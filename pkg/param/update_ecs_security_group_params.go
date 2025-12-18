// Copyright (c) ZStack.io, Inc.

package param

// UpdateEcsSecurityGroupDetailParam UpdateEcsSecurityGroup detail param
type UpdateEcsSecurityGroupDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	Description string `json:"description,omitempty"`
	Name string `json:"name,omitempty"`
}

// UpdateEcsSecurityGroupParam UpdateEcsSecurityGroup request param
type UpdateEcsSecurityGroupParam struct {
	BaseParam
	Params UpdateEcsSecurityGroupDetailParam `json:"params"`
}
