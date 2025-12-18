// Copyright (c) ZStack.io, Inc.

package param

// VerifyLicenseServerDetailParam VerifyLicenseServer详细参数
type VerifyLicenseServerDetailParam struct {
	rest string `json:"appId" validate:"required"` // 必填
	rest string `json:"clientAccessKeyId" validate:"required"` // 必填
	rest string `json:"clientAccessKeySecret" validate:"required"` // 必填
}

// VerifyLicenseServerParam VerifyLicenseServer请求参数
type VerifyLicenseServerParam struct {
	BaseParam
	Params VerifyLicenseServerDetailParam `json:"params"` // 详细参数
}

