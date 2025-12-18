// Copyright (c) ZStack.io, Inc.

package param

// DetachIAM2ProjectFromIAM2OrganizationDetailParam DetachIAM2ProjectFromIAM2Organization详细参数
type DetachIAM2ProjectFromIAM2OrganizationDetailParam struct {
	rest string `json:"projectUuid" validate:"required"` // 必填
}

// DetachIAM2ProjectFromIAM2OrganizationParam DetachIAM2ProjectFromIAM2Organization请求参数
type DetachIAM2ProjectFromIAM2OrganizationParam struct {
	BaseParam
	Params DetachIAM2ProjectFromIAM2OrganizationDetailParam `json:"params"` // 详细参数
}

