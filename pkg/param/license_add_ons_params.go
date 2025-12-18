// Copyright (c) ZStack.io, Inc.

package param

// GetLicenseAddOnsDetailParam GetLicenseAddOns详细参数
type GetLicenseAddOnsDetailParam struct {
}

// GetLicenseAddOnsParam GetLicenseAddOns请求参数
type GetLicenseAddOnsParam struct {
	BaseParam
	Params GetLicenseAddOnsDetailParam `json:"params"` // 详细参数
}

