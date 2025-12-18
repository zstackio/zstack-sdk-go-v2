// Copyright (c) ZStack.io, Inc.

package param

// DetachPolicyFromRoleDetailParam DetachPolicyFromRole详细参数
type DetachPolicyFromRoleDetailParam struct {
	rest string `json:"roleUuid" validate:"required"` // 必填
	rest string `json:"policyUuid" validate:"required"` // 必填
}

// DetachPolicyFromRoleParam DetachPolicyFromRole请求参数
type DetachPolicyFromRoleParam struct {
	BaseParam
	Params DetachPolicyFromRoleDetailParam `json:"params"` // 详细参数
}

