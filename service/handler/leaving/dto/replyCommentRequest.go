package dto

import "encoding/json"

type ReplyCommentRequest struct {
	MsgDataId     int64  `json:"msg_data_id" comment:"群发返回的msg_data_id"`
	Index         int    `json:"index" comment:"多图文时，用来指定第几篇图文，从0开始，不带默认操作该msg_data_id的第一篇图文"`
	UserCommentId int64  `json:"user_comment_id" comment:"评论id"`
	Content       string `json:"content" comment:"回复内容"`
}

func (p *ReplyCommentRequest) ToByte() []byte {
	marshal, err := json.Marshal(p)
	if err != nil {
		return nil
	}
	return marshal
}
