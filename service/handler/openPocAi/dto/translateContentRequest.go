package dto

import "encoding/json"

type TranslateContentRequest struct {
	Content string `json:"content" comment:"源内容(utf8格式，最大600Byte)"`
}

func (p *TranslateContentRequest) ToByte() []byte {
	marshal, err := json.Marshal(p)
	if err != nil {
		return nil
	}
	return marshal
}
