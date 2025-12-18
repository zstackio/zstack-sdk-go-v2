// Copyright (c) ZStack.io, Inc.

package param

// CreateInfoSecSecretResourcePoolDetailParam CreateInfoSecSecretResourcePool详细参数
type CreateInfoSecSecretResourcePoolDetailParam struct {
	rest int `json:"connectionMode" validate:"required"` // 必填
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

// CreateInfoSecSecretResourcePoolParam CreateInfoSecSecretResourcePool请求参数
type CreateInfoSecSecretResourcePoolParam struct {
	BaseParam
	Params CreateInfoSecSecretResourcePoolDetailParam `json:"params"` // 详细参数
}

