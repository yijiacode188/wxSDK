package vo

import "github.com/yijiacode188/wxSDK/service/model/vo"

type GetIndustryResponse struct {
	vo.Base
	PrimaryIndustry struct {
		FirstClass  string `json:"first_class" comment:"一级类目"`
		SecondClass string `json:"second_class" comment:"二级类目"`
	} `json:"primary_industry" comment:"账号设置的主营行业"`
	SecondaryIndustry struct {
		FirstClass  string `json:"first_class" comment:"一级类目"`
		SecondClass string `json:"second_class" comment:"二级类目"`
	} `json:"secondary_industry" comment:"账号设置的副营行业"`
}
