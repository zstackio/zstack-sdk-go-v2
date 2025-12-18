// Copyright (c) ZStack.io, Inc.

package param

// DetachPolicyFromRoleDetailParam DetachPolicyFromRole detail param
type DetachPolicyFromRoleDetailParam struct {
	RoleUuid string `json:"roleUuid" validate:"required"`
	PolicyUuid string `json:"policyUuid" validate:"required"`
}

// DetachPolicyFromRoleParam DetachPolicyFromRole request param
type DetachPolicyFromRoleParam struct {
	BaseParam
	Params DetachPolicyFromRoleDetailParam `json:"params"`
}
