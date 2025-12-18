// Copyright (c) ZStack.io, Inc.

package param

// CheckBaremetalChassisConfigFileDetailParam CheckBaremetalChassisConfigFile详细参数
type CheckBaremetalChassisConfigFileDetailParam struct {
	rest string `json:"baremetalChassisInfo" validate:"required"` // 必填
}

// CheckBaremetalChassisConfigFileParam CheckBaremetalChassisConfigFile请求参数
type CheckBaremetalChassisConfigFileParam struct {
	BaseParam
	Params CheckBaremetalChassisConfigFileDetailParam `json:"params"` // 详细参数
}

