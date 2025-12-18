// Copyright (c) ZStack.io, Inc.

package param

// DetachPolicyFromUserGroupDetailParam DetachPolicyFromUserGroup detail param
type DetachPolicyFromUserGroupDetailParam struct {
	PolicyUuid string `json:"policyUuid" validate:"required"`
	GroupUuid string `json:"groupUuid" validate:"required"`
}

// DetachPolicyFromUserGroupParam DetachPolicyFromUserGroup request param
type DetachPolicyFromUserGroupParam struct {
	BaseParam
	Params DetachPolicyFromUserGroupDetailParam `json:"params"`
}
