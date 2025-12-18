// Copyright (c) ZStack.io, Inc.

package param

// SetImageSecurityLevelDetailParam SetImageSecurityLevel detail param
type SetImageSecurityLevelDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	SecurityLevel string `json:"securityLevel,omitempty"`
}

// SetImageSecurityLevelParam SetImageSecurityLevel request param
type SetImageSecurityLevelParam struct {
	BaseParam
	Params SetImageSecurityLevelDetailParam `json:"params"`
}
