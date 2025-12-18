// Copyright (c) ZStack.io, Inc.

package param

// IsLicenseServerDetailParam IsLicenseServer详细参数
type IsLicenseServerDetailParam struct {
}

// IsLicenseServerParam IsLicenseServer请求参数
type IsLicenseServerParam struct {
	BaseParam
	Params IsLicenseServerDetailParam `json:"params"` // 详细参数
}

