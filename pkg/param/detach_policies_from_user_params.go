// Copyright (c) ZStack.io, Inc.

package param

// DetachPoliciesFromUserDetailParam DetachPoliciesFromUser detail param
type DetachPoliciesFromUserDetailParam struct {
	PolicyUuids []string `json:"policyUuids" validate:"required"`
	UserUuid string `json:"userUuid" validate:"required"`
}

// DetachPoliciesFromUserParam DetachPoliciesFromUser request param
type DetachPoliciesFromUserParam struct {
	BaseParam
	Params DetachPoliciesFromUserDetailParam `json:"params"`
}
