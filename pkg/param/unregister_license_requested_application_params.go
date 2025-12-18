// Copyright (c) ZStack.io, Inc.

package param

// UnregisterLicenseRequestedApplicationDetailParam UnregisterLicenseRequestedApplication详细参数
type UnregisterLicenseRequestedApplicationDetailParam struct {
	rest string `json:"appId" validate:"required"` // 必填
}

// UnregisterLicenseRequestedApplicationParam UnregisterLicenseRequestedApplication请求参数
type UnregisterLicenseRequestedApplicationParam struct {
	BaseParam
	Params UnregisterLicenseRequestedApplicationDetailParam `json:"params"` // 详细参数
}

