// Copyright (c) ZStack.io, Inc.

package param

// GetLicenseUKeyStatusDetailParam GetLicenseUKeyStatus detail param
type GetLicenseUKeyStatusDetailParam struct {
}

// GetLicenseUKeyStatusParam GetLicenseUKeyStatus request param
type GetLicenseUKeyStatusParam struct {
	BaseParam
	Params GetLicenseUKeyStatusDetailParam `json:"params"`
}
