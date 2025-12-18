// Copyright (c) ZStack.io, Inc.

package param

// AddLogConfigurationDetailParam AddLogConfiguration详细参数
type AddLogConfigurationDetailParam struct {
	rest string `json:"name" validate:"required"` // 必填
	rest string `json:"description,omitempty"`
	rest string `json:"type" validate:"required"` // 必填
	rest string `json:"level,omitempty"`
	rest string `json:"configuration" validate:"required"` // 必填
	rest string `json:"resourceUuid,omitempty"`
	rest []string `json:"tagUuids,omitempty"`
}

// AddLogConfigurationParam AddLogConfiguration请求参数
type AddLogConfigurationParam struct {
	BaseParam
	Params AddLogConfigurationDetailParam `json:"params"` // 详细参数
}

