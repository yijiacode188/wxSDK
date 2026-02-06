package notifyTemplate

import (
	"errors"
	"github.com/yijiacode188/wxSDK/service/handler/notifyTemplate/dto"
	"github.com/yijiacode188/wxSDK/service/handler/notifyTemplate/vo"
	"github.com/yijiacode188/wxSDK/utils"
	"net/url"
)

// QueryBlockTmplMsg 查询拦截的模板消息
// 本接口用于查询被拦截的模板消息
// https://developers.weixin.qq.com/doc/service/api/notify/template/api_queryblocktmplmsg.html
func (wx *NotifyTemplate) QueryBlockTmplMsg(body *dto.QueryBlockTmplMsgRequest) (*vo.QueryBlockTmplMsgResponse, *utils.Response, error) {
	token, _, err := wx.GetStableAccessToken(false)
	if err != nil {
		return nil, nil, err
	}
	params := make(url.Values)
	params.Add("access_token", token)
	result, response, err := utils.HttpPost[vo.QueryBlockTmplMsgResponse](&utils.RequestParams{
		Url:    "https://api.weixin.qq.com/wxa/sec/queryblocktmplmsg",
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
