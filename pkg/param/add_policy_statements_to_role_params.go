// Copyright (c) ZStack.io, Inc.

package param

// AddPolicyStatementsToRoleDetailParam AddPolicyStatementsToRole详细参数
type AddPolicyStatementsToRoleDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest []interface{} `json:"statements" validate:"required"` // 必填
}

// AddPolicyStatementsToRoleParam AddPolicyStatementsToRole请求参数
type AddPolicyStatementsToRoleParam struct {
	BaseParam
	Params AddPolicyStatementsToRoleDetailParam `json:"params"` // 详细参数
}

