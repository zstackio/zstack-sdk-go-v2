// Copyright (c) ZStack.io, Inc.

package param

// UpdateRoleDetailParam UpdateRole detail param
type UpdateRoleDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Statements []interface{} `json:"statements,omitempty"`
	PolicyUuids []string `json:"policyUuids,omitempty"`
}

// UpdateRoleParam UpdateRole request param
type UpdateRoleParam struct {
	BaseParam
	Params UpdateRoleDetailParam `json:"params"`
}
