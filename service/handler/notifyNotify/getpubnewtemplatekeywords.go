package notifyNotify

import (
	"errors"
	"github.com/yijiacode188/wxSDK/service/handler/notifyNotify/vo"
	"github.com/yijiacode188/wxSDK/utils"
	"net/url"
)

// GetPubNewTemplateKeywords 获取模板中的关键词
// 该接口用于获取模板标题下的关键词列表
// https://developers.weixin.qq.com/doc/service/api/notify/notify/api_getpubnewtemplatekeywords.html
func (wx *NotifyNotify) GetPubNewTemplateKeywords(tid string) (*vo.GetPubNewTemplateKeywordsResponse, *utils.Response, error) {
	token, _, err := wx.GetStableAccessToken(false)
	if err != nil {
		return nil, nil, err
	}
	params := make(url.Values)
	params.Add("access_token", token)
	params.Add("tid", tid)
	result, response, err := utils.HttpGet[vo.GetPubNewTemplateKeywordsResponse](&utils.RequestParams{
		Url:    "https://api.weixin.qq.com/wxaapi/newtmpl/getpubtemplatekeywords",
		Params: params,
	})
	if err != nil {
		return nil, response, err
	}
	if result.ErrCode != 0 {
		return nil, response, errors.New(result.ErrMsg)
	}
	return &result, response, nil
}
