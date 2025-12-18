// Copyright (c) ZStack.io, Inc.

package param

// UpdateSNSSnmpPlatformDetailParam UpdateSNSSnmpPlatform detail param
type UpdateSNSSnmpPlatformDetailParam struct {
	SnmpAddress string `json:"snmpAddress" validate:"required"`
	SnmpPort int `json:"snmpPort" validate:"required"`
	Uuid string `json:"uuid" validate:"required"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

// UpdateSNSSnmpPlatformParam UpdateSNSSnmpPlatform request param
type UpdateSNSSnmpPlatformParam struct {
	BaseParam
	Params UpdateSNSSnmpPlatformDetailParam `json:"params"`
}
