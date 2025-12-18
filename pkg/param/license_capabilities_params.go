// Copyright (c) ZStack.io, Inc.

package param

// GetLicenseCapabilitiesDetailParam GetLicenseCapabilities详细参数
type GetLicenseCapabilitiesDetailParam struct {
}

// GetLicenseCapabilitiesParam GetLicenseCapabilities请求参数
type GetLicenseCapabilitiesParam struct {
	BaseParam
	Params GetLicenseCapabilitiesDetailParam `json:"params"` // 详细参数
}

