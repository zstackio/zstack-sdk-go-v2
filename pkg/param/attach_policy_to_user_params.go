// Copyright (c) ZStack.io, Inc.

package param

// AttachPolicyToUserDetailParam AttachPolicyToUser detail param
type AttachPolicyToUserDetailParam struct {
	UserUuid string `json:"userUuid" validate:"required"`
	PolicyUuid string `json:"policyUuid" validate:"required"`
}

// AttachPolicyToUserParam AttachPolicyToUser request param
type AttachPolicyToUserParam struct {
	BaseParam
	Params AttachPolicyToUserDetailParam `json:"params"`
}
