package miniprogram

import (
	"errors"
	"github.com/yijiacode188/wxSDK/store"
	"github.com/yijiacode188/wxSDK/types"
)

type wxClient struct {
	AppId  string               `json:"appId"`
	Secret string               `json:"secret"`
	Store  types.StoreInterface `json:"store"`
}

func NewClient(appId, secret string, storeClient ...types.StoreInterface) (*wxClient, error) {
	if appId == "" {
		return nil, errors.New("AppId不能为空")
	}
	if secret == "" {
		return nil, errors.New("secret不能为空")
	}
	client := &wxClient{
		AppId:  appId,
		Secret: secret,
	}
	if len(storeClient) == 0 {
		client.Store = store.NewStorage()
	} else {
		client.Store = storeClient[0]
	}
	return client, nil
}
