// Copyright (c) ZStack.io, Inc.

package param

// SetIAM2ProjectLoginExpiredDetailParam SetIAM2ProjectLoginExpired详细参数
type SetIAM2ProjectLoginExpiredDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"loginExpired" validate:"required"` // 必填
	rest string `json:"loginExpiredAttributeUuid,omitempty"`
}

// SetIAM2ProjectLoginExpiredParam SetIAM2ProjectLoginExpired请求参数
type SetIAM2ProjectLoginExpiredParam struct {
	BaseParam
	Params SetIAM2ProjectLoginExpiredDetailParam `json:"params"` // 详细参数
}

