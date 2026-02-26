package dto

import "encoding/json"

type FreePublishDeleteRequest struct {
	ArticleId string `json:"article_id" comment:"成功发布时返回的 article_id"`
	Index     int    `json:"index" comment:"要删除的文章在图文消息中的位置，第一篇编号为1，该字段不填或填0会删除全部文章"`
}

func (p *FreePublishDeleteRequest) ToByte() []byte {
	marshal, err := json.Marshal(p)
	if err != nil {
		return nil
	}
	return marshal
}
