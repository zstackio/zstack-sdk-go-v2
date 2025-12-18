// Copyright (c) ZStack.io, Inc.

package param

// GetLicenseUKeyStatusDetailParam GetLicenseUKeyStatus详细参数
type GetLicenseUKeyStatusDetailParam struct {
}

// GetLicenseUKeyStatusParam GetLicenseUKeyStatus请求参数
type GetLicenseUKeyStatusParam struct {
	BaseParam
	Params GetLicenseUKeyStatusDetailParam `json:"params"` // 详细参数
}

