package event

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha1"
	"encoding/base64"
	"encoding/xml"
	"errors"
	"fmt"
	"github.com/yijiacode188/wxSDK/service/handler/event/vo"
	"io"
	"regexp"
	"sort"
)

// EventMessage 微信推送的消息
func (wx *Event) EventMessage(bodyData []byte, message MessageEvent) error {
	messageData := &vo.EventMessageResponse{}
	err := xml.Unmarshal(bodyData, messageData)
	if err != nil {
		return err
	}
	if wx.EncodingAESKey == nil {
		//为明文模式
		if !validateAuthPlainText(wx.Token, message.GetTimestamp(), message.GetNonce(), message.GetSignature()) {
			return errors.New("消息验证失败")
		}
	} else {
		//密文模式
		if !validateAuthCiphertext(wx.Token, message.GetTimestamp(), message.GetNonce(), messageData.Encrypt, message.GetMsgSignature()) {
			return errors.New("消息验证失败")
		}
		text, err := decodeCipherText(*wx.EncodingAESKey, messageData.Encrypt)
		if err != nil {
			return err
		}
		xmlStr, err := extractXML(text)
		if err != nil {
			return err
		}
		err = xml.Unmarshal([]byte(xmlStr), messageData)
		if err != nil {
			return err
		}
	}
	//处理回调事件
	switch messageData.MsgType {
	case "text":
		//普通消息
		message.CallBackTextMessage(messageData.ToTextMsg())
	case "image":
		//图片消息
		message.CallBackImageMessage(messageData.ToImageMsg())
	case "voice":
		//语音消息
		message.CallBackVoiceMessage(messageData.ToVoiceMsg())
	case "video":
		//视频消息
		message.CallBackVideoMessage(messageData.ToVideoMsg())
	case "shortvideo":
		//小视频消息
		message.CallBackShortVideoMessage(messageData.ToShortVideoMsg())
	case "location":
		//地理位置消息
		message.CallBackLocationMessage(messageData.ToLocationMsg())
	case "link":
		//链接消息
		message.CallBackLinkMessage(messageData.ToLinkMsg())
	case "event":
		//事件通知
		if messageData.Event == "subscribe" || messageData.Event == "unsubscribe" {
			//订阅或取消订阅
			message.CallBackSubscribeEvent(messageData.ToSubscribeEventMsg())
		}
		if messageData.Event == "SCAN" {
			//扫码
			message.CallBackScanEvent(messageData.ToScanEventMsg())
		}
		if messageData.Event == "LOCATION" {
			//上报地理位置事件
			message.CallBackLocationEvent(messageData.ToLocalEventMsg())
		}
		if messageData.Event == "CLICK" {
			//点击菜单拉取消息时的事件推送
			message.CallBackClickEvent(messageData.ToClickEventMsg())
		}
		if messageData.Event == "VIEW" {
			//点击菜单跳转链接时的事件推送
			message.CallBackViewEvent(messageData.ToViewEventMsg())
		}
	default:
		return errors.New("不支持的消息类型")
	}
	return nil
}

// validateAuthPlainText 明文方式校验参数
func validateAuthPlainText(token, timestamp, nonce, signature string) bool {
	strArr := sort.StringSlice{token, timestamp, nonce}
	sort.Strings(strArr)
	str := ""
	for _, s := range strArr {
		str += s // 拼接字符串
	}
	h := sha1.New()      // 微信用sha1，所以这里使用sha1加密。如果完全都是自己的服务的时候，用md5加密也是没问题的。
	h.Write([]byte(str)) // 转成byte写进digest中。h实际是一个digest类型。
	signatureResult := fmt.Sprintf("%x", h.Sum(nil))
	return signature == signatureResult
}

// validateAuthCiphertext 密文方式校验参数
func validateAuthCiphertext(token, timestamp, nonce, encrypt, msgSignature string) bool {
	strArr := sort.StringSlice{token, timestamp, nonce, encrypt}
	sort.Strings(strArr)
	str := ""
	for _, s := range strArr {
		str += s // 拼接字符串
	}
	h := sha1.New()      // 微信用sha1，所以这里使用sha1加密。如果完全都是自己的服务的时候，用md5加密也是没问题的。
	h.Write([]byte(str)) // 转成byte写进digest中。h实际是一个digest类型。
	signatureResult := fmt.Sprintf("%x", h.Sum(nil))
	return msgSignature == signatureResult
}

// decodeCipherText 密文解密
func decodeCipherText(encodingAESKey, encrypt string) (string, error) {
	aesKey, err := base64.StdEncoding.DecodeString(encodingAESKey + "=")
	if err != nil {
		return "", err
	}

	tmpMsg, err := base64.StdEncoding.DecodeString(encrypt)
	if err != nil {
		return "", err
	}
	result, err := decryptAES128CBC(string(tmpMsg), aesKey)
	if err != nil {
		return "", err
	}
	return result, nil
}

// encodeCipherText 加密
func encodeCipherText(encodingAESKey, text string) (string, error) {
	aesKey, err := base64.StdEncoding.DecodeString(encodingAESKey + "=")
	if err != nil {
		return "", err
	}
	tmpMsg, err := encryptAES128CBC(text, aesKey)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString([]byte(tmpMsg)), nil
}

// DecryptAES128CBC 解密AES-128-CBC加密的字符串
// encryptedStr: 加密字符串（base64编码）
// secretKey: 密钥（16字节）
// 返回解密后的明文
func decryptAES128CBC(encryptedStr string, secretKey []byte) (string, error) {
	// 1. Base64解码
	encryptedData := []byte(encryptedStr)

	// 3. 初始化AES密码块
	block, err := aes.NewCipher(secretKey)
	if err != nil {
		return "", fmt.Errorf("AES初始化失败: %v", err)
	}

	// 4. CBC模式需要IV，通常IV包含在加密数据的前16字节
	iv := encryptedData[:aes.BlockSize]
	ciphertext := encryptedData[aes.BlockSize:]

	// 5. 检查密文长度是否是块大小的倍数
	if len(ciphertext)%aes.BlockSize != 0 {
		return "", fmt.Errorf("密文长度不是块大小的倍数")
	}

	// 6. 创建CBC解密器
	mode := cipher.NewCBCDecrypter(block, iv)

	// 7. 解密（原地解密）
	mode.CryptBlocks(ciphertext, ciphertext)
	// 8. 去除PKCS7填充
	//plaintext, err := pkcs7Unpad(ciphertext, aes.BlockSize)
	//if err != nil {
	//	return "", fmt.Errorf("去除填充失败: %v", err)
	//}

	return string(ciphertext), nil
}

// pkcs7Unpad 去除PKCS7填充
func pkcs7Unpad(data []byte, blockSize int) ([]byte, error) {
	if len(data) == 0 {
		return nil, fmt.Errorf("数据为空")
	}

	if len(data)%blockSize != 0 {
		return nil, fmt.Errorf("数据长度不是块大小的倍数")
	}

	// 获取最后一个字节作为填充长度
	paddingLen := int(data[len(data)-1])

	// 验证填充长度
	if paddingLen <= 0 || paddingLen > blockSize {
		return nil, fmt.Errorf("无效的填充长度")
	}

	// 验证填充字节是否正确
	for i := 0; i < paddingLen; i++ {
		if data[len(data)-1-i] != byte(paddingLen) {
			return nil, fmt.Errorf("填充字节无效")
		}
	}

	// 去除填充
	return data[:len(data)-paddingLen], nil
}
func extractXML(input string) (string, error) {
	if input == "" {
		return "", fmt.Errorf("输入字符串为空")
	}

	// 使用(?s)让.匹配包括换行符的所有字符
	re := regexp.MustCompile(`(?s)<xml>.*?</xml>`)
	xmlPart := re.FindString(input)

	if xmlPart == "" {
		return "", fmt.Errorf("未找到XML部分")
	}

	return xmlPart, nil
}

// EncryptAES128CBC 加密AES-128-CBC模式的字符串
// plaintext: 待加密的明文
// secretKey: 密钥（16字节）
// 返回加密后的base64字符串
func encryptAES128CBC(plaintext string, secretKey []byte) (string, error) {
	// 1. 初始化AES密码块
	block, err := aes.NewCipher(secretKey)
	if err != nil {
		return "", fmt.Errorf("AES初始化失败: %v", err)
	}

	// 2. 生成随机IV（16字节）
	iv := make([]byte, aes.BlockSize)
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", fmt.Errorf("生成IV失败: %v", err)
	}

	// 3. PKCS7填充
	paddedPlaintext := pkcs7Pad([]byte(plaintext), aes.BlockSize)

	// 4. 创建CBC加密器
	mode := cipher.NewCBCEncrypter(block, iv)

	// 5. 加密
	ciphertext := make([]byte, len(paddedPlaintext))
	mode.CryptBlocks(ciphertext, paddedPlaintext)

	// 6. 将IV和密文组合（IV放在前面，与解密对应）
	finalData := append(iv, ciphertext...)

	// 7. Base64编码
	return base64.StdEncoding.EncodeToString(finalData), nil
}

// pkcs7Pad PKCS7填充
func pkcs7Pad(data []byte, blockSize int) []byte {
	if blockSize <= 0 {
		return data
	}

	// 计算需要填充的长度
	paddingLen := blockSize - len(data)%blockSize

	// 创建填充字节（填充字节的值等于填充长度）
	padding := bytes.Repeat([]byte{byte(paddingLen)}, paddingLen)

	// 返回填充后的数据
	return append(data, padding...)
}
