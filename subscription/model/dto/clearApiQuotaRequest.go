package dto

import "encoding/json"

type ClearApiQuotaRequest struct {
	CgiPath string `json:"cgi_path" comment:"api的请求地址"`
}

func (p *ClearApiQuotaRequest) ToByte() []byte {
	marshal, err := json.Marshal(p)
	if err != nil {
		return nil
	}
	return marshal
}
