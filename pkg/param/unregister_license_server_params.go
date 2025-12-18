// Copyright (c) ZStack.io, Inc.

package param

// UnregisterLicenseServerDetailParam UnregisterLicenseServer详细参数
type UnregisterLicenseServerDetailParam struct {
	rest string `json:"clientAuthorizedNodeUuid,omitempty"`
}

// UnregisterLicenseServerParam UnregisterLicenseServer请求参数
type UnregisterLicenseServerParam struct {
	BaseParam
	Params UnregisterLicenseServerDetailParam `json:"params"` // 详细参数
}

