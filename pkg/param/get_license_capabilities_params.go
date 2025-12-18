// Copyright (c) ZStack.io, Inc.

package param

// GetLicenseCapabilitiesDetailParam GetLicenseCapabilities detail param
type GetLicenseCapabilitiesDetailParam struct {
}

// GetLicenseCapabilitiesParam GetLicenseCapabilities request param
type GetLicenseCapabilitiesParam struct {
	BaseParam
	Params GetLicenseCapabilitiesDetailParam `json:"params"`
}
