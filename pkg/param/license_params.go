// Copyright (c) ZStack.io, Inc.

package param

// UpdateLicenseDetailParam UpdateLicense详细参数
type UpdateLicenseDetailParam struct {
	rest string `json:"managementNodeUuid" validate:"required"` // 必填
	rest string `json:"license" validate:"required"` // 必填
	rest string `json:"additionSession,omitempty"`
}

// UpdateLicenseParam UpdateLicense请求参数
type UpdateLicenseParam struct {
	BaseParam
	Params UpdateLicenseDetailParam `json:"params"` // 详细参数
}

