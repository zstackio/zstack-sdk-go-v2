// Copyright (c) ZStack.io, Inc.

package param

// UpdateBaremetalChassisDetailParam UpdateBaremetalChassis detail param
type UpdateBaremetalChassisDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	IpmiAddress string `json:"ipmiAddress,omitempty"`
	IpmiPort int `json:"ipmiPort,omitempty"`
	IpmiUsername string `json:"ipmiUsername,omitempty"`
	IpmiPassword string `json:"ipmiPassword,omitempty"`
}

// UpdateBaremetalChassisParam UpdateBaremetalChassis request param
type UpdateBaremetalChassisParam struct {
	BaseParam
	Params UpdateBaremetalChassisDetailParam `json:"params"`
}
