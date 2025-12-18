// Copyright (c) ZStack.io, Inc.

package param

// GetLicenseInfoDetailParam GetLicenseInfo detail param
type GetLicenseInfoDetailParam struct {
	AdditionSession string `json:"additionSession,omitempty"`
}

// GetLicenseInfoParam GetLicenseInfo request param
type GetLicenseInfoParam struct {
	BaseParam
	Params GetLicenseInfoDetailParam `json:"params"`
}
