// Copyright (c) ZStack.io, Inc.

package param

// RemovePolicyStatementsFromRoleDetailParam RemovePolicyStatementsFromRole详细参数
type RemovePolicyStatementsFromRoleDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest []string `json:"policyStatementUuids" validate:"required"` // 必填
}

// RemovePolicyStatementsFromRoleParam RemovePolicyStatementsFromRole请求参数
type RemovePolicyStatementsFromRoleParam struct {
	BaseParam
	Params RemovePolicyStatementsFromRoleDetailParam `json:"params"` // 详细参数
}

