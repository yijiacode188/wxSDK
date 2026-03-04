package dto

import "encoding/json"

type UserInfoItem struct {
	OpenId string  `json:"openid" comment:"用户的标识，对当前公众号唯一；必须是已关注的用户的 openid"`
	Lang   *string `json:"lang" comment:"国家地区语言版本，zh_CN 简体，zh_TW 繁体，en 英语，默认为zh-CN"`
}
type BatchUserInfoRequest struct {
	UserList []UserInfoItem `json:"user_list" comment:"用户列表"`
}

func (p *BatchUserInfoRequest) ToByte() []byte {
	marshal, err := json.Marshal(p)
	if err != nil {
		return nil
	}
	return marshal
}
