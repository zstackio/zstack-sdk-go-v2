// Copyright (c) ZStack.io, Inc.

package param

// CheckBaremetalChassisConfigFileDetailParam CheckBaremetalChassisConfigFile detail param
type CheckBaremetalChassisConfigFileDetailParam struct {
	BaremetalChassisInfo string `json:"baremetalChassisInfo" validate:"required"`
}

// CheckBaremetalChassisConfigFileParam CheckBaremetalChassisConfigFile request param
type CheckBaremetalChassisConfigFileParam struct {
	BaseParam
	Params CheckBaremetalChassisConfigFileDetailParam `json:"params"`
}
