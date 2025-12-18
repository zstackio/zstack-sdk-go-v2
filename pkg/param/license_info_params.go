// Copyright (c) ZStack.io, Inc.

package param

// GetLicenseInfoDetailParam GetLicenseInfo详细参数
type GetLicenseInfoDetailParam struct {
	rest string `json:"additionSession,omitempty"`
}

// GetLicenseInfoParam GetLicenseInfo请求参数
type GetLicenseInfoParam struct {
	BaseParam
	Params GetLicenseInfoDetailParam `json:"params"` // 详细参数
}

