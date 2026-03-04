package dto

type AuthRequest struct {
	Signature string `json:"signature" comment:"微信加密签名，signature结合了开发者填写的token参数和请求中的timestamp参数、nonce参数。"`
	Timestamp string `json:"timestamp" comment:"时间戳"`
	Nonce     string `json:"nonce" comment:"随机数"`
}
