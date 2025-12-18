// Copyright (c) ZStack.io, Inc.

package param

// UpdateBareMetal2IpmiChassisDetailParam UpdateBareMetal2IpmiChassis detail param
type UpdateBareMetal2IpmiChassisDetailParam struct {
	IpmiAddress string `json:"ipmiAddress,omitempty"`
	IpmiPort int `json:"ipmiPort,omitempty"`
	IpmiUsername string `json:"ipmiUsername,omitempty"`
	IpmiPassword string `json:"ipmiPassword,omitempty"`
	Uuid string `json:"uuid" validate:"required"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

// UpdateBareMetal2IpmiChassisParam UpdateBareMetal2IpmiChassis request param
type UpdateBareMetal2IpmiChassisParam struct {
	BaseParam
	Params UpdateBareMetal2IpmiChassisDetailParam `json:"params"`
}
