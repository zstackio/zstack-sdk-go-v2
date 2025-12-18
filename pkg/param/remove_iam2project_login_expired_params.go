// Copyright (c) ZStack.io, Inc.

package param

// RemoveIAM2ProjectLoginExpiredDetailParam RemoveIAM2ProjectLoginExpired详细参数
type RemoveIAM2ProjectLoginExpiredDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"attributeUuid" validate:"required"` // 必填
}

// RemoveIAM2ProjectLoginExpiredParam RemoveIAM2ProjectLoginExpired请求参数
type RemoveIAM2ProjectLoginExpiredParam struct {
	BaseParam
	Params RemoveIAM2ProjectLoginExpiredDetailParam `json:"params"` // 详细参数
}

