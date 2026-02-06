package dto

import "encoding/json"

type SetIndustryRequest struct {
	IndustryId1 string `json:"industry_id1" comment:"公众号模板消息所属行业编号"`
	IndustryId2 string `json:"industry_id2" comment:"公众号模板消息所属行业编号"`
}

func (p *SetIndustryRequest) ToByte() []byte {
	marshal, err := json.Marshal(p)
	if err != nil {
		return nil
	}
	return marshal
}
