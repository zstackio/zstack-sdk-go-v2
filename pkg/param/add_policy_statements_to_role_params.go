// Copyright (c) ZStack.io, Inc.

package param

// AddPolicyStatementsToRoleDetailParam AddPolicyStatementsToRole detail param
type AddPolicyStatementsToRoleDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	Statements []PolicyStatementParam `json:"statements" validate:"required"`
}

// AddPolicyStatementsToRoleParam AddPolicyStatementsToRole request param
type AddPolicyStatementsToRoleParam struct {
	BaseParam
	Params AddPolicyStatementsToRoleDetailParam `json:"params"`
}
