// Copyright (c) ZStack.io, Inc.

package param

// CreateAiSiNoSecretResourcePoolDetailParam CreateAiSiNoSecretResourcePool详细参数
type CreateAiSiNoSecretResourcePoolDetailParam struct {
	rest string `json:"managementIp,omitempty"`
	rest int `json:"port,omitempty"`
	rest string `json:"route,omitempty"`
	rest string `json:"clientID,omitempty"`
	rest string `json:"clientSecrete,omitempty"`
	rest string `json:"appId,omitempty"`
	rest string `json:"keyNumSM2,omitempty"`
	rest string `json:"keyNumSM4,omitempty"`
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

// CreateAiSiNoSecretResourcePoolParam CreateAiSiNoSecretResourcePool请求参数
type CreateAiSiNoSecretResourcePoolParam struct {
	BaseParam
	Params CreateAiSiNoSecretResourcePoolDetailParam `json:"params"` // 详细参数
}

