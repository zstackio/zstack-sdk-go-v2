// Copyright (c) ZStack.io, Inc.

package param

// DegradeFromLicenseServerDetailParam DegradeFromLicenseServer详细参数
type DegradeFromLicenseServerDetailParam struct {
}

// DegradeFromLicenseServerParam DegradeFromLicenseServer请求参数
type DegradeFromLicenseServerParam struct {
	BaseParam
	Params DegradeFromLicenseServerDetailParam `json:"params"` // 详细参数
}

