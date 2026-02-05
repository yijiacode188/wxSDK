package mpAccessToken

import (
	"errors"
	"github.com/yijiacode188/wxSDK/types"
)

type MPAccessToken struct {
	AppId  string               `json:"appId"`
	Secret string               `json:"secret"`
	Store  types.StoreInterface `json:"store"`
}

func NewMPAccessToken(appId, secret string, storeClient types.StoreInterface) (*MPAccessToken, error) {
	if appId == "" {
		return nil, errors.New("AppId不能为空")
	}
	if secret == "" {
		return nil, errors.New("secret不能为空")
	}
	client := &MPAccessToken{
		AppId:  appId,
		Secret: secret,
		Store:  storeClient,
	}
	return client, nil
}
