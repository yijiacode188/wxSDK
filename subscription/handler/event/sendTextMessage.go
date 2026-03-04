package event

import "github.com/yijiacode188/wxSDK/subscription/handler/event/dto"

// SendTextMessage 发送文字消息
func (wx *Event) SendTextMessage(message *dto.SendTextMessageRequest) ([]byte, error) {
	message.MsgType = "text"
	data, err := message.ToXML()
	if err != nil {
		return nil, err
	}
	if wx.EncodingAESKey == nil {
		//为明文模式
		return data, nil
	} else {
		//为密文模式
		cipher, err := encodeCipherText(*wx.EncodingAESKey, string(data)+wx.AppId)
		if err != nil {
			return nil, err
		}
		return []byte(cipher), nil
	}
}
