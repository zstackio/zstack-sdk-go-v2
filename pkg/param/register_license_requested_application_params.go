// Copyright (c) ZStack.io, Inc.

package param

// RegisterLicenseRequestedApplicationDetailParam RegisterLicenseRequestedApplication详细参数
type RegisterLicenseRequestedApplicationDetailParam struct {
	rest string `json:"licenseRequestCode" validate:"required"` // 必填
	rest string `json:"clientPubKey,omitempty"`
	rest int64 `json:"currentTimeMillis,omitempty"`
}

// RegisterLicenseRequestedApplicationParam RegisterLicenseRequestedApplication请求参数
type RegisterLicenseRequestedApplicationParam struct {
	BaseParam
	Params RegisterLicenseRequestedApplicationDetailParam `json:"params"` // 详细参数
}

