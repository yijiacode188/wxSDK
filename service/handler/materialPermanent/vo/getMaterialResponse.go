package vo

import "github.com/yijiacode188/wxSDK/service/model/vo"

// NewsItem 图文素材，内容
type NewsItem struct {
	Title            string `json:"title" comment:"图文消息的标题"`
	ThumbMediaId     string `json:"thumb_media_id" comment:"图文消息的封面图片素材id（必须是永久mediaID）"`
	ShowCoverPic     int    `json:"show_cover_pic" comment:"是否显示封面，0为false，即不显示，1为true，即显示"`
	Author           string `json:"author" comment:"作者"`
	Digest           string `json:"digest" comment:"图文消息的摘要，仅有单图文消息才有摘要，多图文此处为空"`
	Content          string `json:"content" comment:"图文消息的具体内容，支持HTML标签，必须少于2万字符，小于1M，且此处会去除JS"`
	Url              string `json:"url" comment:"图文页的URL"`
	ContentSourceUrl string `json:"content_source_url" comment:"图文消息的原文地址，即点击“阅读原文”后的URL"`
}
type GetMaterialResponse struct {
	vo.Base
	NewsItem    []NewsItem `json:"news_item" comment:"图文素材，内容"`
	Title       string     `json:"title" comment:"视频素材，标题"`
	Description string     `json:"description" comment:"视频素材，描述"`
	DownUrl     string     `json:"down_url" comment:"视频下载，地址"`
}
