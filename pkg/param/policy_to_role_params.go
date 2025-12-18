// Copyright (c) ZStack.io, Inc.

package param

// AttachPolicyToRoleDetailParam AttachPolicyToRole详细参数
type AttachPolicyToRoleDetailParam struct {
	rest string `json:"roleUuid" validate:"required"` // 必填
	rest string `json:"policyUuid" validate:"required"` // 必填
}

// AttachPolicyToRoleParam AttachPolicyToRole请求参数
type AttachPolicyToRoleParam struct {
	BaseParam
	Params AttachPolicyToRoleDetailParam `json:"params"` // 详细参数
}

