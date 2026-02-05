package miniprogram

import (
	"errors"
	"github.com/yijiacode188/wxSDK/miniprogram/handler/mpAccessToken"
	"github.com/yijiacode188/wxSDK/miniprogram/handler/qrCodeLink"
	"github.com/yijiacode188/wxSDK/miniprogram/handler/userInfo"
	"github.com/yijiacode188/wxSDK/miniprogram/handler/userLogin"
	"github.com/yijiacode188/wxSDK/store"
	"github.com/yijiacode188/wxSDK/types"
)

type wxClient struct {
	*mpAccessToken.MPAccessToken
	*qrCodeLink.QRCodeLink
	*userInfo.UserInfo
	*userLogin.UserLogin
}

// NewClient 初始化小程序客户端
func NewClient(appId, secret string, storeClient ...types.StoreInterface) (*wxClient, error) {
	if appId == "" {
		return nil, errors.New("AppId不能为空")
	}
	if secret == "" {
		return nil, errors.New("secret不能为空")
	}

	var storage types.StoreInterface
	if len(storeClient) == 0 {
		storage = store.NewStorage()
	} else {
		storage = storeClient[0]
	}
	// 基础接口
	mpAccessToken, err := mpAccessToken.NewMPAccessToken(appId, secret, storage)
	if err != nil {
		return nil, err
	}
	// 小程序码与小程序链接
	qrCodeLink, err := qrCodeLink.NewQRCodeLink(mpAccessToken)
	if err != nil {
		return nil, err
	}
	// 用户信息
	userInfo, err := userInfo.NewUserInfo(mpAccessToken)
	if err != nil {
		return nil, err
	}
	// 小程序登录
	userLogin, err := userLogin.NewUserLogin(mpAccessToken)
	if err != nil {
		return nil, err
	}
	client := &wxClient{
		MPAccessToken: mpAccessToken,
		QRCodeLink:    qrCodeLink,
		UserInfo:      userInfo,
		UserLogin:     userLogin,
	}

	return client, nil
}
