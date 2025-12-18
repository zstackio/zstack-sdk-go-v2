// Copyright (c) ZStack.io, Inc.

package param

// UpdateSecurityGroupDetailParam UpdateSecurityGroup detail param
type UpdateSecurityGroupDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

// UpdateSecurityGroupParam UpdateSecurityGroup request param
type UpdateSecurityGroupParam struct {
	BaseParam
	Params UpdateSecurityGroupDetailParam `json:"params"`
}
