package vo

import "github.com/yijiacode188/wxSDK/service/model/vo"

type Category struct {
	Id   int    `json:"id" comment:"类目id，查询公共模板库时需要"`
	Name string `json:"name" comment:"类目的中文名"`
}
type GetCategoryResponse struct {
	vo.Base
	Data []Category `json:"data" comment:"类目列表"`
}
