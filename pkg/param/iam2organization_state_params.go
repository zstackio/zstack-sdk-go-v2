// Copyright (c) ZStack.io, Inc.

package param

// ChangeIAM2OrganizationStateDetailParam ChangeIAM2OrganizationState详细参数
type ChangeIAM2OrganizationStateDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"stateEvent" validate:"required"` // 必填
}

// ChangeIAM2OrganizationStateParam ChangeIAM2OrganizationState请求参数
type ChangeIAM2OrganizationStateParam struct {
	BaseParam
	Params ChangeIAM2OrganizationStateDetailParam `json:"params"` // 详细参数
}

