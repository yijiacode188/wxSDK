package base

import (
	"errors"
	"github.com/yijiacode188/wxSDK/types"
)

type Base struct {
	AppId          string               `json:"appId"`
	Secret         string               `json:"secret"`
	Token          string               `json:"token"`
	EncodingAESKey *string              `json:"encodingAESKey"`
	Store          types.StoreInterface `json:"store"`
}

func NewBase(appId, secret, token string, encodingAESKey *string, storeClient types.StoreInterface) (*Base, error) {
	if appId == "" {
		return nil, errors.New("AppId不能为空")
	}
	if secret == "" {
		return nil, errors.New("secret不能为空")
	}
	client := &Base{
		AppId:          appId,
		Secret:         secret,
		Token:          token,
		Store:          storeClient,
		EncodingAESKey: encodingAESKey,
	}
	return client, nil
}
