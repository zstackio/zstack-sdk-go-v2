// Copyright (c) ZStack.io, Inc.

package param

// AddPolicyStatementsToRoleDetailParam AddPolicyStatementsToRole detail param
type AddPolicyStatementsToRoleDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	Statements []interface{} `json:"statements" validate:"required"`
}

// AddPolicyStatementsToRoleParam AddPolicyStatementsToRole request param
type AddPolicyStatementsToRoleParam struct {
	BaseParam
	Params AddPolicyStatementsToRoleDetailParam `json:"params"`
}
