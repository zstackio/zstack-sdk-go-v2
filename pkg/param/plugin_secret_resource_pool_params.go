// Copyright (c) ZStack.io, Inc.

package param

// CreatePluginSecretResourcePoolDetailParam CreatePluginSecretResourcePool详细参数
type CreatePluginSecretResourcePoolDetailParam struct {
	rest map[string]string `json:"properties,omitempty"`
	rest string `json:"pluginDriverUuid,omitempty"`
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

// CreatePluginSecretResourcePoolParam CreatePluginSecretResourcePool请求参数
type CreatePluginSecretResourcePoolParam struct {
	BaseParam
	Params CreatePluginSecretResourcePoolDetailParam `json:"params"` // 详细参数
}

