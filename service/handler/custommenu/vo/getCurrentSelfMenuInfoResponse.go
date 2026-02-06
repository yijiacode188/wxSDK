package vo

import (
	"github.com/yijiacode188/wxSDK/service/consts"
	"github.com/yijiacode188/wxSDK/service/model/vo"
)

type GetCurrentSelfMenuInfoResponse struct {
	vo.Base
	IsMenuOpen   int `json:"is_menu_open" comment:"菜单是否开启"`
	SelfMenuInfo struct {
		Button []struct {
			Type     consts.ButtonType `json:"type" comment:"菜单的响应动作类型（与 sub_button 互斥）"`
			Name     string            `json:"name" comment:"菜单标题，不超过16个字节，子菜单不超过60个字节"`
			Value    string            `json:"value" comment:"对于不同的菜单类型，value的值意义不同。官网上设置的自定义菜单： Text:保存文字到value； Img、voice：保存mediaID到value； Video：保存视频下载链接到value； News：保存图文消息到news_info，同时保存mediaID到value； View：保存链接到url。 使用API设置的自定义菜单： click、scancode_push、scancode_waitmsg、pic_sysphoto、pic_photo_or_album、 pic_weixin、location_select：保存值到key；view：保存链接到url"`
			Url      string            `json:"url" comment:"对于不同的菜单类型，value的值意义不同。官网上设置的自定义菜单： Text:保存文字到value； Img、voice：保存mediaID到value； Video：保存视频下载链接到value； News：保存图文消息到news_info，同时保存mediaID到value； View：保存链接到url。 使用API设置的自定义菜单： click、scancode_push、scancode_waitmsg、pic_sysphoto、pic_photo_or_album、 pic_weixin、location_select：保存值到key；view：保存链接到url"`
			Key      string            `json:"key" comment:"对于不同的菜单类型，value的值意义不同。官网上设置的自定义菜单： Text:保存文字到value； Img、voice：保存mediaID到value； Video：保存视频下载链接到value； News：保存图文消息到news_info，同时保存mediaID到value； View：保存链接到url。 使用API设置的自定义菜单： click、scancode_push、scancode_waitmsg、pic_sysphoto、pic_photo_or_album、 pic_weixin、location_select：保存值到key；view：保存链接到url"`
			NewsInfo *struct {
				List []struct {
					Title      string `json:"title" comment:"标题"`
					Digest     string `json:"digest" comment:"摘要"`
					Author     string `json:"author" comment:"作者"`
					ShowCover  int    `json:"show_cover" comment:"是否显示封面，0为不显示，1为显示"`
					CoverUrl   string `json:"cover_url" comment:"封面图片的URL"`
					ContentUrl string `json:"content_url" comment:"正文的URL"`
					SourceUrl  string `json:"source_url" comment:"原文的URL，若置空则无查看原文入口"`
				} `json:"list" comment:"图文消息列表"`
			} `json:"news_info" comment:"图文消息的信息"`
		}
	} `json:"selfmenu_info"`
}
