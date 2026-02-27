package vo

import "github.com/yijiacode188/wxSDK/service/model/vo"

type CreateQRCodeResponse struct {
	vo.Base
	Ticket        string `json:"ticket" comment:"获取的二维码ticket，凭借此ticket可以在有效时间内换取二维码。"`
	ExpireSeconds int64  `json:"expire_seconds" comment:"该二维码有效时间，以秒为单位。 最大不超过2592000（即30天）。"`
	Url           string `json:"url" comment:"二维码图片解析后的地址，开发者可根据该地址自行生成需要的二维码图片"`
}
