package notifyMessage

import (
	"errors"
	"github.com/yijiacode188/wxSDK/service/hanlder/notifyMessage/dto"
	"github.com/yijiacode188/wxSDK/service/hanlder/notifyMessage/vo"
	"github.com/yijiacode188/wxSDK/utils"
	"net/url"
)

// Preview 预览消息
// 本接口发送消息给指定用户，在手机端查看消息的样式和排版
// https://developers.weixin.qq.com/doc/service/api/notify/message/api_preview.html
func (wx *NotifyMessage) Preview(body *dto.PreviewRequest) (*utils.Response, error) {
	token, _, err := wx.GetStableAccessToken(false)
	if err != nil {
		return nil, err
	}
	params := make(url.Values)
	params.Add("access_token", token)
	result, response, err := utils.HttpPost[vo.PreviewResponse](&utils.RequestParams{
		Url:    "https://api.weixin.qq.com/cgi-bin/message/mass/preview",
		Params: params,
		Body:   body.ToByte(),
	})
	if err != nil {
		return response, err
	}
	if result.ErrCode != 0 {
		return response, errors.New(result.ErrMsg)
	}
	return response, nil
}
