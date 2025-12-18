// Copyright (c) ZStack.io, Inc.

package param

// CheckBareMetal2IpmiChassisConfigFileDetailParam CheckBareMetal2IpmiChassisConfigFile详细参数
type CheckBareMetal2IpmiChassisConfigFileDetailParam struct {
	rest string `json:"chassisInfo" validate:"required"` // 必填
}

// CheckBareMetal2IpmiChassisConfigFileParam CheckBareMetal2IpmiChassisConfigFile请求参数
type CheckBareMetal2IpmiChassisConfigFileParam struct {
	BaseParam
	Params CheckBareMetal2IpmiChassisConfigFileDetailParam `json:"params"` // 详细参数
}

