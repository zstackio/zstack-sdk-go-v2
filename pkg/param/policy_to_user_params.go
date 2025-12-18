// Copyright (c) ZStack.io, Inc.

package param

// AttachPolicyToUserDetailParam AttachPolicyToUser详细参数
type AttachPolicyToUserDetailParam struct {
	rest string `json:"userUuid" validate:"required"` // 必填
	rest string `json:"policyUuid" validate:"required"` // 必填
}

// AttachPolicyToUserParam AttachPolicyToUser请求参数
type AttachPolicyToUserParam struct {
	BaseParam
	Params AttachPolicyToUserDetailParam `json:"params"` // 详细参数
}

