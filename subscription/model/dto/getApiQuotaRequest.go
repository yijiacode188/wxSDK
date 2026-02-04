package dto

import "encoding/json"

type GetApiQuotaRequest struct {
	CgiPath string `json:"cgi_path" comment:"api的请求地址，例如\"/cgi-bin/message/custom/send\""`
}

func (p *GetApiQuotaRequest) ToByte() []byte {
	marshal, err := json.Marshal(p)
	if err != nil {
		return nil
	}
	return marshal
}
