// Copyright (c) ZStack.io, Inc.

package param

// GetLicenseAddOnsDetailParam GetLicenseAddOns detail param
type GetLicenseAddOnsDetailParam struct {
}

// GetLicenseAddOnsParam GetLicenseAddOns request param
type GetLicenseAddOnsParam struct {
	BaseParam
	Params GetLicenseAddOnsDetailParam `json:"params"`
}
