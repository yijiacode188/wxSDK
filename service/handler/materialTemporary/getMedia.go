package materialTemporary

import (
	"errors"
	"github.com/yijiacode188/wxSDK/service/handler/materialTemporary/vo"
	"github.com/yijiacode188/wxSDK/utils"
	"net/url"
)

// GetMedia 获取临时素材
// 本接口用于获取临时素材（即下载临时的多媒体文件）。
// https://developers.weixin.qq.com/doc/service/api/material/temporary/api_getmedia.html
func (wx *MaterialTemporary) GetMedia(mediaId string) (string, *utils.Response, error) {
	token, _, err := wx.GetStableAccessToken(false)
	if err != nil {
		return "", nil, err
	}
	params := make(url.Values)
	params.Add("access_token", token)
	params.Add("media_id", mediaId)
	result, response, err := utils.HttpGet[vo.GetMediaResponse](&utils.RequestParams{
		Url:    "https://api.weixin.qq.com/cgi-bin/media/get",
		Params: params,
	})

	if err != nil {
		return "", response, err
	}
	if result.ErrCode != 0 {
		return "", response, errors.New(result.ErrMsg)
	}
	return result.VideoUrl, response, nil
}
