// Copyright (c) ZStack.io, Inc.

package param

// CreateBaremetalInstanceDetailParam CreateBaremetalInstance详细参数
type CreateBaremetalInstanceDetailParam struct {
	rest string `json:"name" validate:"required"` // 必填
	rest string `json:"description,omitempty"`
	rest string `json:"chassisUuid" validate:"required"` // 必填
	rest string `json:"imageUuid" validate:"required"` // 必填
	rest string `json:"templateUuid,omitempty"`
	rest string `json:"username,omitempty"`
	rest string `json:"password" validate:"required"` // 必填
	rest map[string]string `json:"nicCfgs,omitempty"`
	rest map[string]string `json:"bondingCfgs,omitempty"`
	rest map[string]string `json:"customConfigurations,omitempty"`
	rest string `json:"strategy,omitempty"`
	rest string `json:"resourceUuid,omitempty"`
	rest []string `json:"tagUuids,omitempty"`
}

// CreateBaremetalInstanceParam CreateBaremetalInstance请求参数
type CreateBaremetalInstanceParam struct {
	BaseParam
	Params CreateBaremetalInstanceDetailParam `json:"params"` // 详细参数
}

