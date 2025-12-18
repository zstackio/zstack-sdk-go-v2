// Copyright (c) ZStack.io, Inc.

package param

// CreateHaiTaiSecretResourcePoolDetailParam CreateHaiTaiSecretResourcePool详细参数
type CreateHaiTaiSecretResourcePoolDetailParam struct {
	rest string `json:"managementIp,omitempty"`
	rest int `json:"port,omitempty"`
	rest string `json:"realm,omitempty"`
	rest string `json:"name" validate:"required"` // 必填
	rest string `json:"description,omitempty"`
	rest string `json:"model,omitempty"`
	rest string `json:"ability,omitempty"`
	rest string `json:"type" validate:"required"` // 必填
	rest int `json:"heartbeatInterval" validate:"required"` // 必填
	rest string `json:"zoneUuid" validate:"required"` // 必填
	rest string `json:"resourceUuid,omitempty"`
	rest []string `json:"tagUuids,omitempty"`
}

// CreateHaiTaiSecretResourcePoolParam CreateHaiTaiSecretResourcePool请求参数
type CreateHaiTaiSecretResourcePoolParam struct {
	BaseParam
	Params CreateHaiTaiSecretResourcePoolDetailParam `json:"params"` // 详细参数
}

