package notifyMessage

import (
	"errors"
	"github.com/yijiacode188/wxSDK/service/handler/notifyMessage/dto"
	"github.com/yijiacode188/wxSDK/service/handler/notifyMessage/vo"
	"github.com/yijiacode188/wxSDK/utils"
	"net/url"
)

// UploadNewsMsg 上传图文消息素材
// 本接口用于上传图文消息，该能力已更新为草稿箱
// https://developers.weixin.qq.com/doc/service/api/notify/message/api_uploadnewsmsg.html
func (wx *NotifyMessage) UploadNewsMsg(body *dto.UploadNewsMsgRequest) (*vo.UploadNewsMsgResponse, *utils.Response, error) {
	token, _, err := wx.GetStableAccessToken(false)
	if err != nil {
		return nil, nil, err
	}
	params := make(url.Values)
	params.Add("access_token", token)
	result, response, err := utils.HttpPost[vo.UploadNewsMsgResponse](&utils.RequestParams{
		Url:    "https://api.weixin.qq.com/cgi-bin/media/uploadnews",
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
