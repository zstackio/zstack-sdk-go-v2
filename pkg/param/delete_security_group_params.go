// Copyright (c) ZStack.io, Inc.

package param

// DeleteSecurityGroupDetailParam DeleteSecurityGroup detail param
type DeleteSecurityGroupDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// DeleteSecurityGroupParam DeleteSecurityGroup request param
type DeleteSecurityGroupParam struct {
	BaseParam
	Params DeleteSecurityGroupDetailParam `json:"params"`
}
