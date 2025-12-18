// Copyright (c) ZStack.io, Inc.

package param

// AttachSecurityGroupToL3NetworkDetailParam AttachSecurityGroupToL3Network详细参数
type AttachSecurityGroupToL3NetworkDetailParam struct {
	rest string `json:"securityGroupUuid" validate:"required"` // 必填
	rest string `json:"l3NetworkUuid" validate:"required"` // 必填
}

// AttachSecurityGroupToL3NetworkParam AttachSecurityGroupToL3Network请求参数
type AttachSecurityGroupToL3NetworkParam struct {
	BaseParam
	Params AttachSecurityGroupToL3NetworkDetailParam `json:"params"` // 详细参数
}

