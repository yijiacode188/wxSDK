package event

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha1"
	"encoding/base64"
	"encoding/xml"
	"errors"
	"fmt"
	"github.com/yijiacode188/wxSDK/service/handler/event/dto"
	"github.com/yijiacode188/wxSDK/service/handler/event/vo"
	"regexp"
	"sort"
)

// EventMessage 微信推送的消息
func (wx *Event) EventMessage(query *dto.EventMessageRequestQuery, body *dto.EventMessageBody) (*vo.EventMessageResponse, error) {
	out := &vo.EventMessageResponse{}
	if wx.EncodingAESKey == nil {
		//为明文模式
		if !validateAuthPlainText(wx.Token, query.Timestamp, query.Nonce, query.Signature) {
			return nil, errors.New("消息验证失败")
		}
		out = &body.EventMessageResponse
	} else {
		//密文模式
		if !validateAuthCiphertext(wx.Token, query.Timestamp, query.Nonce, body.Encrypt, query.MsgSignature) {
			return nil, errors.New("消息验证失败")
		}
		text, err := decodeCipherText(*wx.EncodingAESKey, body.Encrypt)
		if err != nil {
			return nil, err
		}
		xmlStr, err := extractXML(text)
		if err != nil {
			return nil, err
		}
		err = xml.Unmarshal([]byte(xmlStr), out)
		if err != nil {
			return nil, err
		}
	}
	return out, nil
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
	plaintext, err := pkcs7Unpad(ciphertext, aes.BlockSize)
	if err != nil {
		return "", fmt.Errorf("去除填充失败: %v", err)
	}

	return string(plaintext), nil
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
