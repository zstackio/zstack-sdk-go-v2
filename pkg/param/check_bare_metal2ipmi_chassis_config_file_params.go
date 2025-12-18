// Copyright (c) ZStack.io, Inc.

package param

// CheckBareMetal2IpmiChassisConfigFileDetailParam CheckBareMetal2IpmiChassisConfigFile detail param
type CheckBareMetal2IpmiChassisConfigFileDetailParam struct {
	ChassisInfo string `json:"chassisInfo" validate:"required"`
}

// CheckBareMetal2IpmiChassisConfigFileParam CheckBareMetal2IpmiChassisConfigFile request param
type CheckBareMetal2IpmiChassisConfigFileParam struct {
	BaseParam
	Params CheckBareMetal2IpmiChassisConfigFileDetailParam `json:"params"`
}
