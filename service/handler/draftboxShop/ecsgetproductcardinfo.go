package draftboxShop

import (
	"errors"
	"github.com/yijiacode188/wxSDK/service/handler/draftboxShop/dto"
	"github.com/yijiacode188/wxSDK/service/handler/draftboxShop/vo"
	"github.com/yijiacode188/wxSDK/utils"
	"net/url"
)

// EcsGetProductCardInfo 获取商品卡片的DOM结构
// 本接口用于获取在文章中插入商品卡片所需的DOM结构
// 注：不同文章类型支持的卡片类型范围不同
// 图片消息：支持小卡、文字链接、条卡
// 图文消息：支持大卡、小卡、文字链接
// https://developers.weixin.qq.com/doc/service/api/draftbox/shop/api_ecsgetproductcardinfo.html
func (wx *DraftBoxShop) EcsGetProductCardInfo(body *dto.EcsGetProductCardInfoRequest) (*vo.EcsGetProductCardInfoResponse, *utils.Response, error) {
	token, _, err := wx.GetStableAccessToken(false)
	if err != nil {
		return nil, nil, err
	}
	params := make(url.Values)
	params.Add("access_token", token)
	result, response, err := utils.HttpPost[vo.EcsGetProductCardInfoResponse](&utils.RequestParams{
		Url:    "https://api.weixin.qq.com/channels/ec/service/product/getcardinfo",
		Params: params,
		Body:   body.ToByte(),
	})
	if err != nil {
		return nil, response, err
	}
	if result.ErrCode != 0 {
		return nil, response, errors.New(result.ErrMsg)
	}
	return &result, response, nil
}
