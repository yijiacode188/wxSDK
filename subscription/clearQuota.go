package subscription

import (
	"errors"
	"github.com/yijiacode188/wxSDK/subscription/model/dto"
	"github.com/yijiacode188/wxSDK/subscription/model/vo"
	"github.com/yijiacode188/wxSDK/utils"
	"net/url"
)

func (wx *wxClient) ClearQuota(body *dto.ClearQuotaRequest) error {
	token, err := wx.GetStableAccessToken(false)
	if err != nil {
		return err
	}
	params := make(url.Values)
	params.Add("access_token", token)
	result, _, err := utils.HttpPost[vo.ClearQuotaResponse](&utils.RequestParams{
		Url:    "https://api.weixin.qq.com/cgi-bin/clear_quota",
		Params: params,
		Body:   body.ToByte(),
	})
	if err != nil {
		return err
	}
	if result.ErrCode != 0 {
		return errors.New(result.ErrMsg)
	}
	return nil
}
