package dto

import "encoding/json"

type ArticleType string

const (
	ArticleTypeNews    ArticleType = "news"    //图文消息
	ArticleTypeNewsPic ArticleType = "newspic" //图片消息
)

type EcsGetProductCardInfoRequest struct {
	ProductId   string      `json:"product_id" comment:"商品id"`
	ArticleType ArticleType `json:"article_type" comment:"文章类型，当前支持：图片消息(newspic)、图文消息(news)"`
	CardType    int         `json:"card_type" comment:"卡片类型，当前支持：大卡(0)、小卡(1)、文字链接(2)、条卡(3)"`
}

func (p *EcsGetProductCardInfoRequest) ToByte() []byte {
	marshal, err := json.Marshal(p)
	if err != nil {
		return nil
	}
	return marshal
}
