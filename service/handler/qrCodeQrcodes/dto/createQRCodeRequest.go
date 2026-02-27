package dto

import "encoding/json"

type ActionName string

const (
	ActionNameQrScene         ActionName = "QR_SCENE"           // 临时整型
	ActionNameQrStrScene      ActionName = "QR_STR_SCENE"       // 临时字符串
	ActionNameQrLimitScene    ActionName = "QR_LIMIT_SCENE"     //永久整型
	ActionNameQrLimitStrScene ActionName = "QR_LIMIT_STR_SCENE" //永久字符串
)

type Scene struct {
	SceneId  int64  `json:"scene_id" comment:"场景值ID，临时二维码时为32位非0整型，永久二维码时最大值为100000（目前参数只支持1--100000）"`
	SceneStr string `json:"scene_str" comment:"场景值ID（字符串形式的ID），字符串类型，长度限制为1到64"`
}
type ActionInfo struct {
	Scene Scene `json:"scene" comment:"场景信息"`
}
type CreateQRCodeRequest struct {
	ExpireSeconds int64      `json:"expire_seconds" comment:"二维码有效时间（秒），最大2592000，仅临时二维码需要"`
	ActionName    ActionName `json:"action_name" comment:"二维码类型：QR_SCENE(临时整型)/QR_STR_SCENE(临时字符串)/QR_LIMIT_SCENE(永久整型)/QR_LIMIT_STR_SCENE(永久字符串)"`
	ActionInfo    ActionInfo `json:"action_info" comment:"二维码详细信息"`
}

func (p *CreateQRCodeRequest) ToByte() []byte {
	marshal, err := json.Marshal(p)
	if err != nil {
		return nil
	}
	return marshal
}
