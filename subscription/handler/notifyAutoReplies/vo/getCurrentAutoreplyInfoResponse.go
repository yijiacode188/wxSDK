package vo

import "github.com/yijiacode188/wxSDK/subscription/model/vo"

type AddFriendAutoReplyInfo struct {
	Type    string `json:"type" comment:"自动回复的类型。关注后自动回复和消息自动回复的类型仅支持文本（text）、图片（img）、语音（voice）、视频（video），关键词自动回复则还多了图文消息（news）"`
	Content string `json:"content" comment:"对于文本类型，content是文本内容，对于图文、图片、语音、视频类型，content是mediaID"`
}
type MessageDefaultAutoReplyInfo struct {
	Type    string `json:"type" comment:"自动回复的类型。关注后自动回复和消息自动回复的类型仅支持文本（text）、图片（img）、语音（voice）、视频（video），关键词自动回复则还多了图文消息（news）"`
	Content string `json:"content" comment:"对于文本类型，content是文本内容，对于图文、图片、语音、视频类型，content是mediaID"`
}
type KeywordAutoReplyInfo struct {
	RuleName        string `json:"rule_name" comment:"规则名称"`
	CreateTime      int    `json:"create_time" comment:"创建时间"`
	ReplyMode       string `json:"reply_mode" comment:"回复模式，reply_all代表全部回复，random_one代表随机回复其中一条"`
	KeywordListInfo []struct {
		Type     string `json:"type" comment:"自动回复的类型。关注后自动回复和消息自动回复的类型仅支持文本（text）、图片（img）、语音（voice）、视频（video），关键词自动回复则还多了图文消息（news）"`
		Content  string `json:"content" comment:"对于文本类型，content是文本内容，对于图文、图片、语音、视频类型，content是mediaID"`
		MathMode string `json:"match_mode" comment:"匹配模式，contain代表消息中含有该关键词即可，equal表示消息内容必须和关键词严格相同"`
	} `json:"keyword_list_info" comment:"匹配的关键词列表"`
	ReplyListInfo []struct {
		Type     string `json:"type" comment:"自动回复的类型。关注后自动回复和消息自动回复的类型仅支持文本（text）、图片（img）、语音（voice）、视频（video），关键词自动回复则还多了图文消息（news）"`
		Content  string `json:"content" comment:"对于文本类型，content是文本内容，对于图文、图片、语音、视频类型，content是mediaID"`
		NewsInfo struct {
			List []struct {
				Title      string `json:"title" comment:"图文消息的标题"`
				Digest     string `json:"digest" comment:"摘要"`
				Author     string `json:"author" comment:"作者"`
				ShowCover  int    `json:"show_cover" comment:"是否显示封面，0为不显示，1为显示"`
				CoverUrl   string `json:"cover_url" comment:"封面图片的URL"`
				ContentUrl string `json:"content_url" comment:"正文的URL"`
				SourceUrl  string `json:"source_url" comment:"原文的URL，若置空则无查看原文入口"`
			} `json:"list" comment:"图文消息列表"`
		} `json:"news_info" comment:"图文消息的信息"`
	}
}
type GetCurrentAutoReplyInfoResponse struct {
	vo.Base
	IsAddFriendReplyOpen        int                         `json:"is_add_friend_reply_open" comment:"关注后自动回复是否开启(0-未开启 1-开启)"`
	IsAutoReplyOpen             int                         `json:"is_autoreply_open" comment:"消息自动回复是否开启(0-未开启 1-开启)"`
	AddFriendAutoReplyInfo      AddFriendAutoReplyInfo      `json:"add_friend_autoreply_info" comment:"关注后自动回复的信息"`
	MessageDefaultAutoReplyInfo MessageDefaultAutoReplyInfo `json:"message_default_autoreply_info" comment:"消息自动回复的信息"`
	KeywordAutoReplyInfo        KeywordAutoReplyInfo        `json:"keyword_autoreply_info" comment:"关键词自动回复的信息"`
}
