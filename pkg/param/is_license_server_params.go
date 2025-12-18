// Copyright (c) ZStack.io, Inc.

package param

// IsLicenseServerDetailParam IsLicenseServer detail param
type IsLicenseServerDetailParam struct {
}

// IsLicenseServerParam IsLicenseServer request param
type IsLicenseServerParam struct {
	BaseParam
	Params IsLicenseServerDetailParam `json:"params"`
}
