package vo

import "github.com/yijiacode188/wxSDK/service/model/vo"

type TemplateItem struct {
	TemplateId      string `json:"template_id" comment:"模板ID"`
	Title           string `json:"title" comment:"模板标题"`
	PrimaryIndustry string `json:"primary_industry" comment:"模板所属行业的一级行业"`
	DeputyIndustry  string `json:"deputy_industry" comment:"模板所属行业的二级行业"`
	Content         string `json:"content" comment:"模板内容"`
	Example         string `json:"example" comment:"模板示例"`
}
type GetAllTemplatesResponse struct {
	vo.Base
	TemplateList []TemplateItem `json:"template_list" comment:"模版列表"`
}
