// Copyright (c) ZStack.io, Inc.

package param

// VerifyLicenseServerDetailParam VerifyLicenseServer detail param
type VerifyLicenseServerDetailParam struct {
	AppId string `json:"appId" validate:"required"`
	ClientAccessKeyId string `json:"clientAccessKeyId" validate:"required"`
	ClientAccessKeySecret string `json:"clientAccessKeySecret" validate:"required"`
}

// VerifyLicenseServerParam VerifyLicenseServer request param
type VerifyLicenseServerParam struct {
	BaseParam
	Params VerifyLicenseServerDetailParam `json:"params"`
}
