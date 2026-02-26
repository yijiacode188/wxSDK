package vo

import "github.com/yijiacode188/wxSDK/service/model/vo"

type FreePublishGetArticleResponse struct {
	vo.Base
	NewsItem []struct {
		Title              string `json:"title" comment:"标题"`
		Author             string `json:"author" comment:"作者"`
		Digest             string `json:"digest" comment:"图文消息的摘要，仅有单图文消息才有摘要，多图文此处为空。如果本字段为没有填写，则默认抓取正文前54个字。"`
		Content            string `json:"content" comment:"图文消息的具体内容，支持HTML标签，必须少于2万字符，小于1M，且此处会去除JS"`
		ContentSourceUrl   string `json:"content_source_url" comment:"图文消息的原文地址，即点击“阅读原文”后的URL"`
		ThumbMediaId       string `json:"thumb_media_id" comment:"图文消息的封面图片素材id（必须是永久MediaID）"`
		ThumbUrl           string `json:"thumb_url" comment:"图文消息的封面图片URL"`
		NeedOpenComment    int    `json:"need_open_comment" comment:"是否打开评论，0不打开(默认)，1打开"`
		OnlyFansCanComment int    `json:"only_fans_can_comment" comment:"是否粉丝才可评论，0所有人可评论(默认)，1粉丝才可评论"`
		Url                string `json:"url" comment:"草稿的临时链接"`
		IsDelete           bool   `json:"is_delete" comment:"该图文是否被删除"`
	} `json:"news_item" comment:"图文信息集合"`
}
