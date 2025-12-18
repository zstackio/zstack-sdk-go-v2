// Copyright (c) ZStack.io, Inc.

package param

// CreateBareMetal2IpmiChassisHardwareInfoDetailParam CreateBareMetal2IpmiChassisHardwareInfo detail param
type CreateBareMetal2IpmiChassisHardwareInfoDetailParam struct {
	IpmiAddress string `json:"ipmiAddress" validate:"required"`
	IpmiPort int `json:"ipmiPort" validate:"required"`
	HardwareInfo string `json:"hardwareInfo" validate:"required"`
	ConvertInfo string `json:"convertInfo,omitempty"`
}

// CreateBareMetal2IpmiChassisHardwareInfoParam CreateBareMetal2IpmiChassisHardwareInfo request param
type CreateBareMetal2IpmiChassisHardwareInfoParam struct {
	BaseParam
	Params CreateBareMetal2IpmiChassisHardwareInfoDetailParam `json:"params"`
}
