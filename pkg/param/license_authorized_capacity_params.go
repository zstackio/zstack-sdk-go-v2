// Copyright (c) ZStack.io, Inc.

package param

// GetLicenseAuthorizedCapacityDetailParam GetLicenseAuthorizedCapacity详细参数
type GetLicenseAuthorizedCapacityDetailParam struct {
	rest string `json:"clientAuthorizedNodeUuid,omitempty"`
	rest bool `json:"showServerCapacity,omitempty"`
}

// GetLicenseAuthorizedCapacityParam GetLicenseAuthorizedCapacity请求参数
type GetLicenseAuthorizedCapacityParam struct {
	BaseParam
	Params GetLicenseAuthorizedCapacityDetailParam `json:"params"` // 详细参数
}

