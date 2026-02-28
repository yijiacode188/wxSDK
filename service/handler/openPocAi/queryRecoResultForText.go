package openPocAi

import (
	"errors"
	"github.com/yijiacode188/wxSDK/service/consts"
	"github.com/yijiacode188/wxSDK/service/handler/openPocAi/vo"
	"github.com/yijiacode188/wxSDK/utils"
	"net/url"
)

// queryReCoResultForText 获取语音识别结果
// 本接口用于查询语音转文字结果
// https://developers.weixin.qq.com/doc/service/api/openpoc/ai/api_queryrecoresultfortext.html
func (wx *OpenPocAi) QueryReCoResultForText(voiceId string, lang *consts.TranslateLang) (*vo.QueryReCoResultForTextResponse, *utils.Response, error) {
	token, _, err := wx.GetStableAccessToken(false)
	if err != nil {
		return nil, nil, err
	}
	params := make(url.Values)
	params.Add("access_token", token)
	params.Add("voice_id", voiceId)
	if lang != nil {
		params.Add("lang", string(*lang))
	}

	result, response, err := utils.HttpPost[vo.QueryReCoResultForTextResponse](&utils.RequestParams{
		Url:    "https://api.weixin.qq.com/cgi-bin/media/voice/queryrecoresultfortext",
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
