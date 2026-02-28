package openPocAi

import (
	"errors"
	"github.com/yijiacode188/wxSDK/service/consts"
	"github.com/yijiacode188/wxSDK/service/handler/openPocAi/dto"
	"github.com/yijiacode188/wxSDK/service/handler/openPocAi/vo"
	"github.com/yijiacode188/wxSDK/utils"
	"net/url"
)

// TranslateContent 微信翻译
// 本接口用于文本内容翻译
// https://developers.weixin.qq.com/doc/service/api/openpoc/ai/api_translatecontent.html
func (wx *OpenPocAi) TranslateContent(langFrom, toFrom consts.TranslateLang, body *dto.TranslateContentRequest) (*vo.TranslateContentResponse, *utils.Response, error) {
	token, _, err := wx.GetStableAccessToken(false)
	if err != nil {
		return nil, nil, err
	}
	params := make(url.Values)
	params.Add("access_token", token)
	params.Add("lfrom", string(langFrom))
	params.Add("lto", string(toFrom))
	result, response, err := utils.HttpPost[vo.TranslateContentResponse](&utils.RequestParams{
		Url:    "https://api.weixin.qq.com/cgi-bin/media/voice/translatecontent",
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
