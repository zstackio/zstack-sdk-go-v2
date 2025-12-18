// Copyright (c) ZStack.io, Inc.

package param

// DetachPoliciesFromUserDetailParam DetachPoliciesFromUser详细参数
type DetachPoliciesFromUserDetailParam struct {
	rest []string `json:"policyUuids" validate:"required"` // 必填
	rest string `json:"userUuid" validate:"required"` // 必填
}

// DetachPoliciesFromUserParam DetachPoliciesFromUser请求参数
type DetachPoliciesFromUserParam struct {
	BaseParam
	Params DetachPoliciesFromUserDetailParam `json:"params"` // 详细参数
}

