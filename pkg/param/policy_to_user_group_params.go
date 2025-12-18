// Copyright (c) ZStack.io, Inc.

package param

// AttachPolicyToUserGroupDetailParam AttachPolicyToUserGroup详细参数
type AttachPolicyToUserGroupDetailParam struct {
	rest string `json:"policyUuid" validate:"required"` // 必填
	rest string `json:"groupUuid" validate:"required"` // 必填
}

// AttachPolicyToUserGroupParam AttachPolicyToUserGroup请求参数
type AttachPolicyToUserGroupParam struct {
	BaseParam
	Params AttachPolicyToUserGroupDetailParam `json:"params"` // 详细参数
}

