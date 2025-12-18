// Copyright (c) ZStack.io, Inc.

package param

// UpdateLogConfigurationDetailParam UpdateLogConfiguration详细参数
type UpdateLogConfigurationDetailParam struct {
	rest int64 `json:"configId" validate:"required"` // 必填
	rest string `json:"name,omitempty"`
	rest string `json:"description,omitempty"`
}

// UpdateLogConfigurationParam UpdateLogConfiguration请求参数
type UpdateLogConfigurationParam struct {
	BaseParam
	Params UpdateLogConfigurationDetailParam `json:"params"` // 详细参数
}

