// Copyright (c) ZStack.io, Inc.

package param

// AttachIAM2ProjectToIAM2OrganizationDetailParam AttachIAM2ProjectToIAM2Organization详细参数
type AttachIAM2ProjectToIAM2OrganizationDetailParam struct {
	rest string `json:"projectUuid" validate:"required"` // 必填
	rest string `json:"organizationUuid" validate:"required"` // 必填
}

// AttachIAM2ProjectToIAM2OrganizationParam AttachIAM2ProjectToIAM2Organization请求参数
type AttachIAM2ProjectToIAM2OrganizationParam struct {
	BaseParam
	Params AttachIAM2ProjectToIAM2OrganizationDetailParam `json:"params"` // 详细参数
}

