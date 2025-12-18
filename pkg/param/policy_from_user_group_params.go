// Copyright (c) ZStack.io, Inc.

package param

// DetachPolicyFromUserGroupDetailParam DetachPolicyFromUserGroup详细参数
type DetachPolicyFromUserGroupDetailParam struct {
	rest string `json:"policyUuid" validate:"required"` // 必填
	rest string `json:"groupUuid" validate:"required"` // 必填
}

// DetachPolicyFromUserGroupParam DetachPolicyFromUserGroup请求参数
type DetachPolicyFromUserGroupParam struct {
	BaseParam
	Params DetachPolicyFromUserGroupDetailParam `json:"params"` // 详细参数
}

