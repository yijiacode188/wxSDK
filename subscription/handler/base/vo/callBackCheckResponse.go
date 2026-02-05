package vo

import (
	"github.com/yijiacode188/wxSDK/subscription/model/vo"
)

type CallBackCheckResponse struct {
	vo.Base
	DNS []struct {
		IP           string `json:"ip" comment:"解析出来的ip"`
		RealOperator string `json:"real_operator" comment:"ip对应的运营商"`
	} `json:"dns"`
	PING []struct {
		IP           string `json:"ip" comment:"ping的ip，执行命令为ping ip –c 1-w 1 -q"`
		FromOperator string `json:"from_operator" comment:"ping的源头的运营商，由请求中的check_operator控制"`
		PackageLoss  string `json:"package_loss" comment:"ping的丢包率，0%表示无丢包，100%表示全部丢包。因为目前仅发送一个ping包，因此取值仅有0%或者100%两种可能。"`
	} `json:"ping"`
}
