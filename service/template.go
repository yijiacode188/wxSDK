package service

import (
	"errors"
	"github.com/yijiacode188/wxSDK/service/hanlder/apimanage"
	"github.com/yijiacode188/wxSDK/service/hanlder/base"
	"github.com/yijiacode188/wxSDK/service/hanlder/custommenu"
	"github.com/yijiacode188/wxSDK/service/hanlder/notifyAutoReplies"
	"github.com/yijiacode188/wxSDK/service/hanlder/notifyMessage"
	"github.com/yijiacode188/wxSDK/service/hanlder/notifyNotify"
	"github.com/yijiacode188/wxSDK/service/hanlder/notifySubscribe"
	"github.com/yijiacode188/wxSDK/service/hanlder/notifyTemplate"
	"github.com/yijiacode188/wxSDK/store"
	"github.com/yijiacode188/wxSDK/types"
)

type WxClient struct {
	*base.Base
	*apimanage.ApiManager
	*custommenu.CustomMenu
	*notifyMessage.NotifyMessage
	*notifyTemplate.NotifyTemplate
	*notifyNotify.NotifyNotify
	*notifySubscribe.NotifySubscribe
	*notifyAutoReplies.NotifyAutoReplies
}

// NewClient 初始化服务号客户端
func NewClient(appId, secret string, storeClient ...types.StoreInterface) (*WxClient, error) {
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
	baseContext, err := base.NewBase(appId, secret, storage)
	if err != nil {
		return nil, err
	}

	// openApi管理接口
	apiManagerContext, err := apimanage.NewApiManager(baseContext)
	if err != nil {
		return nil, err
	}

	// 自定义菜单接口
	customMenuContext, err := custommenu.NewCustomMenu(baseContext)
	if err != nil {
		return nil, err
	}

	//基础消息与订阅通知 群发消息
	notifyMessage, err := notifyMessage.NewNotifyMessage(baseContext)
	if err != nil {
		return nil, err
	}

	//基础消息与订阅通知 模版消息
	notifyTemplate, err := notifyTemplate.NewNotifyTemplate(baseContext)
	if err != nil {
		return nil, err
	}

	//基础消息与订阅通知 订阅通知
	notifyNotify, err := notifyNotify.NewNotifyNotify(baseContext)
	if err != nil {
		return nil, err
	}

	//基础消息与订阅通知 一次性订阅消息
	notifySubscribe, err := notifySubscribe.NewNotifyNotify(baseContext)
	if err != nil {
		return nil, err
	}

	//基础消息与订阅通知 自动回复
	notifyAutoReplies, err := notifyAutoReplies.NewNotifyAutoReplies(baseContext)
	if err != nil {
		return nil, err
	}
	wxClient := &WxClient{
		Base:              baseContext,
		ApiManager:        apiManagerContext,
		CustomMenu:        customMenuContext,
		NotifyMessage:     notifyMessage,
		NotifyTemplate:    notifyTemplate,
		NotifyNotify:      notifyNotify,
		NotifySubscribe:   notifySubscribe,
		NotifyAutoReplies: notifyAutoReplies,
	}

	return wxClient, nil
}
