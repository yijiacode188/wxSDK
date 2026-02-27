package vo

import "github.com/yijiacode188/wxSDK/service/model/vo"

type QrCodeJumpGetResponse struct {
	vo.Base
	RuleList []struct {
		Prefix      string   `json:"prefix" comment:"二维码规则，说明，服务号二维码规则已过滤不展示。"`
		Path        string   `json:"path" comment:"小程序功能页面。"`
		State       int      `json:"state" comment:"发布标志位，1 表示未发布，2 表示已发布。"`
		OpenVersion int      `json:"open_version" comment:"1表示开发版（配置只对开发者生效）；2表示体验版（配置对管理员、体验者生效)；3表示正式版（配置对开发者、管理员和体验者生效）"`
		DebugUrl    []string `json:"debug_url" comment:"测试链接（选填）可填写不多于 5 个用于测试的二维码完整链接，此链接必须符合已填写的二维码规则。获取“扫普通二维码打开小程序”已设置的二维码规则才返回这个参数。"`
	} `json:"rule_list" comment:"二维码规则详情列表"`
	QrCodeJumpOpen     int   `json:"qrcodejump_open" comment:"是否已经打开二维码跳转链接设置"`
	ListSize           int64 `json:"list_size" comment:"二维码规则数量"`
	QrCodeJumpPubQuota int64 `json:"qrcodejump_pub_quota" comment:"本月还可发布的次数"`
	TotalCount         int64 `json:"total_count" comment:"二维码规则总数据量，用于分页查询"`
}
