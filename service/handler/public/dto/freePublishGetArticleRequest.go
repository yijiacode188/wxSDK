package dto

import "encoding/json"

type FreePublishGetArticleRequest struct {
	ArticleId string `json:"article_id" comment:"要获取的草稿的article_id"`
}

func (p *FreePublishGetArticleRequest) ToByte() []byte {
	marshal, err := json.Marshal(p)
	if err != nil {
		return nil
	}
	return marshal
}
