// Copyright (c) ZStack.io, Inc.

package param

// ReloadLicenseDetailParam ReloadLicense detail param
type ReloadLicenseDetailParam struct {
	ManagementNodeUuids []string `json:"managementNodeUuids,omitempty"`
	AdditionSession string `json:"additionSession,omitempty"`
}

// ReloadLicenseParam ReloadLicense request param
type ReloadLicenseParam struct {
	BaseParam
	Params ReloadLicenseDetailParam `json:"params"`
}
