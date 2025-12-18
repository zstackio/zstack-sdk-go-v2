// Copyright (c) ZStack.io, Inc.

package param

// GetGuestOsMetadataDetailParam GetGuestOsMetadata detail param
type GetGuestOsMetadataDetailParam struct {
}

// GetGuestOsMetadataParam GetGuestOsMetadata request param
type GetGuestOsMetadataParam struct {
	BaseParam
	Params GetGuestOsMetadataDetailParam `json:"params"`
}
