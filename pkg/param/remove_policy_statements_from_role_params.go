// Copyright (c) ZStack.io, Inc.

package param

// RemovePolicyStatementsFromRoleDetailParam RemovePolicyStatementsFromRole detail param
type RemovePolicyStatementsFromRoleDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	PolicyStatementUuids []string `json:"policyStatementUuids" validate:"required"`
}

// RemovePolicyStatementsFromRoleParam RemovePolicyStatementsFromRole request param
type RemovePolicyStatementsFromRoleParam struct {
	BaseParam
	Params RemovePolicyStatementsFromRoleDetailParam `json:"params"`
}
