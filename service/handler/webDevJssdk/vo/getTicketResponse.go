package vo

import "github.com/yijiacode188/wxSDK/service/model/vo"

type GetTicketResponse struct {
	vo.Base
	Ticket    string `json:"ticket" comment:"临时票据"`
	ExpiresIn int64  `json:"expires_in" comment:"有效期（秒）"`
}
