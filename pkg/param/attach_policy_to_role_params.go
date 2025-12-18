// Copyright (c) ZStack.io, Inc.

package param

// AttachPolicyToRoleDetailParam AttachPolicyToRole detail param
type AttachPolicyToRoleDetailParam struct {
	RoleUuid string `json:"roleUuid" validate:"required"`
	PolicyUuid string `json:"policyUuid" validate:"required"`
}

// AttachPolicyToRoleParam AttachPolicyToRole request param
type AttachPolicyToRoleParam struct {
	BaseParam
	Params AttachPolicyToRoleDetailParam `json:"params"`
}
