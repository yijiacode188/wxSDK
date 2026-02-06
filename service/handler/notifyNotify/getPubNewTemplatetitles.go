package notifyNotify

import (
	"errors"
	"fmt"
	"github.com/yijiacode188/wxSDK/service/handler/notifyNotify/vo"
	"github.com/yijiacode188/wxSDK/utils"
	"net/url"
)

// GetPubNewTemplateTitles 获取类目下的公共模板
// 该接口用于获取帐号所属类目下的公共模板，可从中选用模板使用
// https://developers.weixin.qq.com/doc/service/api/notify/notify/api_getpubnewtemplatetitles.html
func (wx *NotifyNotify) GetPubNewTemplateTitles(ids string, start, limit int) (*vo.GetPubNewTemplateTitlesResponse, *utils.Response, error) {
	token, _, err := wx.GetStableAccessToken(false)
	if err != nil {
		return nil, nil, err
	}
	params := make(url.Values)
	params.Add("access_token", token)
	params.Add("ids", ids)
	params.Add("start", fmt.Sprintf("%d", start))
	params.Add("limit", fmt.Sprintf("%d", limit))
	result, response, err := utils.HttpGet[vo.GetPubNewTemplateTitlesResponse](&utils.RequestParams{
		Url:    "https://api.weixin.qq.com/wxaapi/newtmpl/getpubtemplatetitles",
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
