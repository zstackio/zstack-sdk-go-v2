// Copyright (c) ZStack.io, Inc.

package param

// DetachSecurityGroupFromL3NetworkDetailParam DetachSecurityGroupFromL3Network详细参数
type DetachSecurityGroupFromL3NetworkDetailParam struct {
	rest string `json:"securityGroupUuid" validate:"required"` // 必填
	rest string `json:"l3NetworkUuid" validate:"required"` // 必填
}

// DetachSecurityGroupFromL3NetworkParam DetachSecurityGroupFromL3Network请求参数
type DetachSecurityGroupFromL3NetworkParam struct {
	BaseParam
	Params DetachSecurityGroupFromL3NetworkDetailParam `json:"params"` // 详细参数
}

