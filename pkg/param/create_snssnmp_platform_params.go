// Copyright (c) ZStack.io, Inc.

package param

// CreateSNSSnmpPlatformDetailParam CreateSNSSnmpPlatform detail param
type CreateSNSSnmpPlatformDetailParam struct {
	SnmpAddress string `json:"snmpAddress" validate:"required"`
	SnmpPort int `json:"snmpPort" validate:"required"`
	Name string `json:"name" validate:"required"`
	Description string `json:"description,omitempty"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateSNSSnmpPlatformParam CreateSNSSnmpPlatform request param
type CreateSNSSnmpPlatformParam struct {
	BaseParam
	Params CreateSNSSnmpPlatformDetailParam `json:"params"`
}
