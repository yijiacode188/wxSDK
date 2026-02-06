package notifyMessage

import (
	"errors"
	"github.com/yijiacode188/wxSDK/service/handler/notifyMessage/dto"
	"github.com/yijiacode188/wxSDK/service/handler/notifyMessage/vo"
	"github.com/yijiacode188/wxSDK/utils"
	"net/url"
)

// DeleteMassMsg 删除群发消息
// 群发之后，随时可以通过该接口删除群发。
// https://developers.weixin.qq.com/doc/service/api/notify/message/api_deletemassmsg.html
func (wx *NotifyMessage) DeleteMassMsg(body *dto.DeleteMassMsgRequest) (*utils.Response, error) {
	token, _, err := wx.GetStableAccessToken(false)
	if err != nil {
		return nil, err
	}
	params := make(url.Values)
	params.Add("access_token", token)
	result, response, err := utils.HttpPost[vo.DeleteMassMsgResponse](&utils.RequestParams{
		Url:    "https://api.weixin.qq.com/cgi-bin/message/mass/delete",
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
