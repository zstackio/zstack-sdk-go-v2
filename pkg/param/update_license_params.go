// Copyright (c) ZStack.io, Inc.

package param

// UpdateLicenseDetailParam UpdateLicense detail param
type UpdateLicenseDetailParam struct {
	ManagementNodeUuid string `json:"managementNodeUuid" validate:"required"`
	License string `json:"license" validate:"required"`
	AdditionSession string `json:"additionSession,omitempty"`
}

// UpdateLicenseParam UpdateLicense request param
type UpdateLicenseParam struct {
	BaseParam
	Params UpdateLicenseDetailParam `json:"params"`
}
