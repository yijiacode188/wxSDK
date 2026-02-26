package vo

import "github.com/yijiacode188/wxSDK/service/model/vo"

type DraftItem struct {
	MediaId    string `json:"media_id" comment:"图文消息的id"`
	UpdateTime int64  `json:"update_time" comment:"图文消息更新时间"`
	Content    *struct {
		NewsItem []struct {
			ArticleType        string `json:"article_type" comment:"文章类型，分别有图文消息（news）、图片消息（newspic），不填默认为图文消息（news）"`
			Title              string `json:"title" comment:"标题"`
			Author             string `json:"author" comment:"作者"`
			Digest             string `json:"digest" comment:"图文消息的摘要，仅有单图文消息才有摘要，多图文此处为空。如果本字段为没有填写，则默认抓取正文前54个字。"`
			Content            string `json:"content" comment:"图文消息的具体内容，支持HTML标签，必须少于2万字符，小于1M，且此处会去除JS,涉及图片url必须来源"`
			ContentSourceUrl   string `json:"content_source_url" comment:"图文消息的原文地址，即点击“阅读原文”后的URL"`
			ThumbMediaId       string `json:"thumb_media_id" comment:"图文消息的封面图片素材id（必须是永久MediaID）"`
			NeedOpenComment    int    `json:"need_open_comment" comment:"是否打开评论，0不打开(默认)，1打开"`
			OnlyFansCanComment int    `json:"only_fans_can_comment" comment:"是否粉丝才可评论，0所有人可评论(默认)，1粉丝才可评论"`
			ImageInfo          *struct {
				ImageList []struct {
					ImageMediaId string `json:"image_media_id" comment:"图片消息里的图片素材id（必须是永久MediaID）"`
				} `json:"image_list" comment:"图片列表"`
			} `json:"image_info" comment:"图片消息里的图片相关信息，图片数量最多为20张，首张图片即为封面图"`
			ProductInfo *struct {
				FooterProductInfo struct {
					ProductKey string `json:"product_key" comment:"商品key"`
				} `json:"footer_product_info" comment:"文末插入商品相关信息"`
			} `json:"product_info" comment:"商品信息"`
		} `json:"news_item" comment:"图文内容列表"`
	} `json:"content" comment:"图文消息内容"`
}
type DraftBatchGetResponse struct {
	vo.Base
	TotalCount int64       `json:"total_count" comment:"草稿素材的总数"`
	ItemCount  int64       `json:"item_count" comment:"本次调用获取的素材的数量"`
	Item       []DraftItem `json:"item" comment:"素材列表"`
}
