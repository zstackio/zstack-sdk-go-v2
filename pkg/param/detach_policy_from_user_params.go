// Copyright (c) ZStack.io, Inc.

package param

// DetachPolicyFromUserDetailParam DetachPolicyFromUser detail param
type DetachPolicyFromUserDetailParam struct {
	PolicyUuid string `json:"policyUuid" validate:"required"`
	UserUuid string `json:"userUuid" validate:"required"`
}

// DetachPolicyFromUserParam DetachPolicyFromUser request param
type DetachPolicyFromUserParam struct {
	BaseParam
	Params DetachPolicyFromUserDetailParam `json:"params"`
}
