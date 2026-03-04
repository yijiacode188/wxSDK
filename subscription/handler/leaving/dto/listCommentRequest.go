package dto

import "encoding/json"

type ListCommentRequest struct {
	MsgDataId int64 `json:"msg_data_id" comment:"群发返回的msg_data_id"`
	Index     int   `json:"index" comment:"多图文时，用来指定第几篇图文，从0开始，不带默认返回该msg_data_id的第一篇图文"`
	Begin     int   `json:"begin" comment:"起始位置"`
	Count     int64 `json:"count" comment:"获取数目（>=50会被拒绝）"`
	Type      int   `json:"type" comment:"type=0 普通评论&精选评论 type=1 普通评论 type=2 精选评论"`
}

func (p *ListCommentRequest) ToByte() []byte {
	marshal, err := json.Marshal(p)
	if err != nil {
		return nil
	}
	return marshal
}
