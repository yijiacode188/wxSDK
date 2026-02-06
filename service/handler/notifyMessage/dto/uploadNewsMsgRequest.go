package dto

import "encoding/json"

type Article struct {
	Title              string `json:"title" comment:"图文消息的标题"`
	Author             string `json:"author" comment:"图文消息的作者"`
	ThumbMediaId       string `json:"thumb_media_id" comment:"图文消息缩略图的media_id，可以在素材管理-新增临时素材中获得"`
	Content            string `json:"content" comment:"图文消息页面的内容，支持HTML标签。具备微信支付权限的公众号，可以使用a标签，其他公众号不能使用。"`
	ContentSourceUrl   string `json:"content_source_url" comment:"在图文消息页面点击“阅读原文”后的页面，受安全限制，如需跳转Appstore，可以使用itun.es或appsto.re的短链服务，并在短链后增加 #wechat_redirect 后缀"`
	Digest             string `json:"digest" comment:"图文消息的描述，如本字段为空，则默认抓取正文前64个字"`
	ShowCoverPic       int    `json:"show_cover_pic" comment:"是否显示封面，1为显示，0为不显示"`
	NeedOpenComment    int    `json:"need_open_comment" comment:"是否打开评论，0不打开，1打开"`
	OnlyFansCanComment int    `json:"only_fans_can_comment" comment:"是否粉丝才可评论，0所有人可评论，1粉丝才可评论"`
}
type UploadNewsMsgRequest struct {
	Articles []Article `json:"articles" comment:"图文消息，一个图文消息支持1到8条图文"`
}

func (p *UploadNewsMsgRequest) ToByte() []byte {
	marshal, err := json.Marshal(p)
	if err != nil {
		return nil
	}
	return marshal
}
