package vo

import "github.com/yijiacode188/wxSDK/service/model/vo"

type EcsGetProductCardInfoResponse struct {
	vo.Base
	ProductKey string `json:"product_key" comment:"商品 key，部分文章类型插入商品卡片需要该 key"`
	Dom        string `json:"DOM" comment:"商品卡DOM结构，多数文章类型插入商品卡片需要DOM结构"`
}
