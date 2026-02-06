package dto

import "encoding/json"

type DeleteMassMsgRequest struct {
	MsgId      int    `json:"msg_id" comment:"发送出去的消息ID"`
	ArticleIdx int    `json:"article_idx" comment:"要删除的文章在图文消息中的位置，第一篇编号为1，该字段不填或填0会删除全部文章"`
	Url        string `json:"url" comment:"要删除的文章url，当msg_id未指定时该参数才生效"`
}

func (p *DeleteMassMsgRequest) ToByte() []byte {
	marshal, err := json.Marshal(p)
	if err != nil {
		return nil
	}
	return marshal
}
