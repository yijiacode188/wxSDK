package vo

import "github.com/yijiacode188/wxSDK/service/model/vo"

type SnsUserInfoResponse struct {
	vo.Base
	OpenID     string   `json:"openid" comment:"用户的唯一标识"`
	Nickname   string   `json:"nickname" comment:"用户昵称"`
	Sex        int      `json:"sex" comment:"用户的性别，值为1时是男性，值为2时是女性，值为0时是未知"`
	HeadImgUrl string   `json:"headimgurl" comment:"用户头像，最后一个数值代表正方形头像大小（有0、46、64、96、132数值可选，0代表640*640正方形头像），用户没有头像时该项为空。若用户更换头像，原有头像URL将失效。"`
	UnionId    string   `json:"unionid" comment:"只有在管理员将公众号绑定到微信开放平台账号后，才会出现该字段。"`
	Province   string   `json:"province" comment:"用户个人资料填写的省份"`
	City       string   `json:"city" comment:"普通用户个人资料填写的城市"`
	Country    string   `json:"country" comment:"国家，如中国为CN"`
	Privilege  []string `json:"privilege" comment:"用户特权信息，json 数组，如微信沃卡用户为（chinaunicom）"`
}
