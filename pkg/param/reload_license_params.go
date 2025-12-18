// Copyright (c) ZStack.io, Inc.

package param

// ReloadLicenseDetailParam ReloadLicense详细参数
type ReloadLicenseDetailParam struct {
	rest []string `json:"managementNodeUuids,omitempty"`
	rest string `json:"additionSession,omitempty"`
}

// ReloadLicenseParam ReloadLicense请求参数
type ReloadLicenseParam struct {
	BaseParam
	Params ReloadLicenseDetailParam `json:"params"` // 详细参数
}

