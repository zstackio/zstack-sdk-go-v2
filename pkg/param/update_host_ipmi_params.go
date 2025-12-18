// Copyright (c) ZStack.io, Inc.

package param

// UpdateHostIpmiDetailParam UpdateHostIpmi detail param
type UpdateHostIpmiDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	IpmiAddress string `json:"ipmiAddress,omitempty"`
	IpmiUsername string `json:"ipmiUsername,omitempty"`
	IpmiPassword string `json:"ipmiPassword,omitempty"`
	IpmiPort int `json:"ipmiPort,omitempty"`
}

// UpdateHostIpmiParam UpdateHostIpmi request param
type UpdateHostIpmiParam struct {
	BaseParam
	Params UpdateHostIpmiDetailParam `json:"params"`
}
