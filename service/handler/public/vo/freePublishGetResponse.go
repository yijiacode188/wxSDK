package vo

import "github.com/yijiacode188/wxSDK/service/model/vo"

type FreePublishGetResponse struct {
	vo.Base
	PublishId     string `json:"publish_id" comment:"发布任务id"`
	PublishStatus int    `json:"publish_status" comment:"发布状态(0:成功,1:发布中,2:原创失败,3:常规失败,4:平台审核不通过,5:成功后用户删除所有文章,6:成功后系统封禁所有文章)"`
	ArticleId     string `json:"article_id" comment:"成功时的图文article_id"`
	ArticleDetail *struct {
		Count int64 `json:"count" comment:"当发布状态为0时（即成功）时，返回文章数量"`
		Item  []struct {
			Idx        int64  `json:"idx" comment:"文章对应的编号"`
			ArticleUrl string `json:"article_url" comment:"图文的永久链接"`
		} `json:"item" comment:"当发布状态为0时（即成功）时，返回文章详情"`
	} `json:"article_detail" comment:"文章详情"`
	FailIdx []int64 `json:"fail_idx" comment:"失败文章编号"`
}
