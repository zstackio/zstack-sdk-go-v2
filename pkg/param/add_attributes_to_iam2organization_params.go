// Copyright (c) ZStack.io, Inc.

package param

// AddAttributesToIAM2OrganizationDetailParam AddAttributesToIAM2Organization详细参数
type AddAttributesToIAM2OrganizationDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest []interface{} `json:"attributes" validate:"required"` // 必填
}

// AddAttributesToIAM2OrganizationParam AddAttributesToIAM2Organization请求参数
type AddAttributesToIAM2OrganizationParam struct {
	BaseParam
	Params AddAttributesToIAM2OrganizationDetailParam `json:"params"` // 详细参数
}

