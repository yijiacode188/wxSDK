package vo

import "github.com/yijiacode188/wxSDK/service/model/vo"

type MaterialItem struct {
	MediaId string `json:"media_id" comment:"消息id"`
	Content struct {
		NewsItem []NewsItem `json:"news_item" comment:"图文消息内的1篇或多篇文章"`
	} `json:"content" comment:"图文消息，内容"`
	UpdateTime int64  `json:"update_time" comment:"更新日期"`
	Name       string `json:"name" comment:"图片、语音、视频素材的名字"`
	Url        string `json:"url" comment:"图片、语音、视频素材URL"`
}
type BatchGetMaterialResponse struct {
	vo.Base
	Item       []MaterialItem `json:"item" comment:"多个图文消息"`
	TotalCount int64          `json:"total_count" comment:"该类型的素材的总数"`
	ItemCount  int64          `json:"item_count" comment:"本次调用获取的素材的数量"`
}
