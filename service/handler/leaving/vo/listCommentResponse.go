package vo

import "github.com/yijiacode188/wxSDK/subscription/model/vo"

type ListCommentResponse struct {
	vo.Base
	Total   int64 `json:"total" comment:"评论总数"`
	Comment []struct {
		UserCommentId int64  `json:"user_comment_id" comment:"用户评论id"`
		CreateTime    int64  `json:"create_time" comment:"评论时间 "`
		Content       string `json:"content" comment:"评论内容 "`
		CommentType   int    `json:"comment_type" comment:"是否精选评论，0为即非精选，1为true，即精选"`
		OpenId        string `json:"openid" comment:"openid，用户如果用非微信身份评论，不返回openid"`
		Reply         struct {
			Content    string `json:"content" comment:"回复内容"`
			CreateTime int64  `json:"create_time" comment:"回复时间 "`
		} `json:"reply" comment:"回复信息"`
	} `json:"comment" comment:"评论列表"`
}
