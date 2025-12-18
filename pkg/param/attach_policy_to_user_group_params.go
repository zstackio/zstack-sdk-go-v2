// Copyright (c) ZStack.io, Inc.

package param

// AttachPolicyToUserGroupDetailParam AttachPolicyToUserGroup detail param
type AttachPolicyToUserGroupDetailParam struct {
	PolicyUuid string `json:"policyUuid" validate:"required"`
	GroupUuid string `json:"groupUuid" validate:"required"`
}

// AttachPolicyToUserGroupParam AttachPolicyToUserGroup request param
type AttachPolicyToUserGroupParam struct {
	BaseParam
	Params AttachPolicyToUserGroupDetailParam `json:"params"`
}
