// Copyright (c) ZStack.io, Inc.

package param

// DetachPolicyFromUserDetailParam DetachPolicyFromUser详细参数
type DetachPolicyFromUserDetailParam struct {
	rest string `json:"policyUuid" validate:"required"` // 必填
	rest string `json:"userUuid" validate:"required"` // 必填
}

// DetachPolicyFromUserParam DetachPolicyFromUser请求参数
type DetachPolicyFromUserParam struct {
	BaseParam
	Params DetachPolicyFromUserDetailParam `json:"params"` // 详细参数
}

