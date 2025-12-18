// Copyright (c) ZStack.io, Inc.

package param

// SetVmSecurityLevelDetailParam SetVmSecurityLevel detail param
type SetVmSecurityLevelDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	SecurityLevel string `json:"securityLevel,omitempty"`
}

// SetVmSecurityLevelParam SetVmSecurityLevel request param
type SetVmSecurityLevelParam struct {
	BaseParam
	Params SetVmSecurityLevelDetailParam `json:"params"`
}
