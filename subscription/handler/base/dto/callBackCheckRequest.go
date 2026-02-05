package dto

import "encoding/json"

type CallBackCheckRequest struct {
	Action        string `json:"action" comment:"检测动作：dns(域名解析)/ping(ping检测)/all(全部)"`
	CheckOperator string `json:"check_operator" comment:"检测运营商：CHINANET(电信)/UNICOM(联通)/CAP(腾讯)/DEFAULT(自动)"`
}

func (p *CallBackCheckRequest) ToByte() []byte {
	if p.Action == "" {
		p.Action = "all"
	}
	if p.CheckOperator == "" {
		p.CheckOperator = "DEFAULT"
	}
	marshal, err := json.Marshal(p)
	if err != nil {
		return nil
	}
	return marshal
}
