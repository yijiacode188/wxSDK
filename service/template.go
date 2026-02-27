package service

import (
	"errors"
	"github.com/yijiacode188/wxSDK/service/handler/apimanage"
	"github.com/yijiacode188/wxSDK/service/handler/base"
	"github.com/yijiacode188/wxSDK/service/handler/custommenu"
	"github.com/yijiacode188/wxSDK/service/handler/draftboxDraftmanage"
	"github.com/yijiacode188/wxSDK/service/handler/draftboxShop"
	"github.com/yijiacode188/wxSDK/service/handler/leaving"
	"github.com/yijiacode188/wxSDK/service/handler/materialPermanent"
	"github.com/yijiacode188/wxSDK/service/handler/materialTemporary"
	"github.com/yijiacode188/wxSDK/service/handler/notifyAutoReplies"
	"github.com/yijiacode188/wxSDK/service/handler/notifyMessage"
	"github.com/yijiacode188/wxSDK/service/handler/notifyNotify"
	"github.com/yijiacode188/wxSDK/service/handler/notifySubscribe"
	"github.com/yijiacode188/wxSDK/service/handler/notifyTemplate"
	"github.com/yijiacode188/wxSDK/service/handler/public"
	"github.com/yijiacode188/wxSDK/service/handler/qrCodeQrCodeJump"
	"github.com/yijiacode188/wxSDK/service/handler/qrCodeQrcodes"
	"github.com/yijiacode188/wxSDK/service/handler/userManagerTag"
	"github.com/yijiacode188/wxSDK/service/handler/userManagerUserInfo"
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
	*materialPermanent.MaterialPermanent
	*materialTemporary.MaterialTemporary
	*draftboxDraftmanage.DraftBoxDraftManage
	*draftboxShop.DraftBoxShop
	*leaving.Leaving
	*public.Public
	*userManagerTag.UserManagerTag
	*userManagerUserInfo.UserManagerUserInfo
	*qrCodeQrCodeJump.QrCodeQrCodeJump
	*qrCodeQrcodes.QrCodeQrCodes
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

	//素材管理 永久素材
	materialPermanent, err := materialPermanent.NewMaterialPermanent(baseContext)
	if err != nil {
		return nil, err
	}

	//素材管理 临时素材
	materialTemporary, err := materialTemporary.NewMaterialTemporary(baseContext)
	if err != nil {
		return nil, err
	}
	// 草稿管理和商品卡片 草稿管理
	draftBoxDraftManage, err := draftboxDraftmanage.NewDraftBoxDraftManage(baseContext)
	if err != nil {
		return nil, err
	}
	// 草稿管理和商品卡片 商品卡片
	draftBoxShop, err := draftboxShop.NewDraftBoxShop(baseContext)
	if err != nil {
		return nil, err
	}
	// 留言管理
	leaving, err := leaving.NewLeaving(baseContext)
	if err != nil {
		return nil, err
	}

	// 发布能力
	public, err := public.NewPublic(baseContext)
	if err != nil {
		return nil, err
	}

	// 用户管理 标签管理
	userManagerTag, err := userManagerTag.NewUserManagerTag(baseContext)
	if err != nil {
		return nil, err
	}

	// 用户管理 用户信息
	userManagerUserInfo, err := userManagerUserInfo.NewUserManagerUserInfo(baseContext)
	if err != nil {
		return nil, err
	}

	// 服务号二维码 扫二维码打开小程序
	qrCodeQrCodeJump, err := qrCodeQrCodeJump.NewQrCodeQrCodeJump(baseContext)
	if err != nil {
		return nil, err
	}

	// 服务号二维码 带参二维码
	qrCodeQrcodes, err := qrCodeQrcodes.NewQrCodeQrCodes(baseContext)
	if err != nil {
		return nil, err
	}

	wxClient := &WxClient{
		Base:                baseContext,
		ApiManager:          apiManagerContext,
		CustomMenu:          customMenuContext,
		NotifyMessage:       notifyMessage,
		NotifyTemplate:      notifyTemplate,
		NotifyNotify:        notifyNotify,
		NotifySubscribe:     notifySubscribe,
		NotifyAutoReplies:   notifyAutoReplies,
		MaterialPermanent:   materialPermanent,
		MaterialTemporary:   materialTemporary,
		DraftBoxDraftManage: draftBoxDraftManage,
		DraftBoxShop:        draftBoxShop,
		Leaving:             leaving,
		Public:              public,
		UserManagerTag:      userManagerTag,
		UserManagerUserInfo: userManagerUserInfo,
		QrCodeQrCodeJump:    qrCodeQrCodeJump,
		QrCodeQrCodes:       qrCodeQrcodes,
	}

	return wxClient, nil
}
