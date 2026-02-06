package dto

import "encoding/json"

type AddTemplateRequest struct {
	TemplateIdShort string   `json:"template_id_short" comment:"模板库中模板的编号，有“TM**”和“OPENTMTM**”等形式,对于类目模板，为纯数字ID"`
	KeywordNameList []string `json:"keyword_name_list" comment:"选用的类目模板的关键词,按顺序传入,如果为空，或者关键词不在模板库中，会返回40246错误码"`
}

func (p *AddTemplateRequest) ToByte() []byte {
	marshal, err := json.Marshal(p)
	if err != nil {
		return nil
	}
	return marshal
}
