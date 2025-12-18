// Copyright (c) ZStack.io, Inc.

package param

// RefreshGuestOsMetadataDetailParam RefreshGuestOsMetadata detail param
type RefreshGuestOsMetadataDetailParam struct {
}

// RefreshGuestOsMetadataParam RefreshGuestOsMetadata request param
type RefreshGuestOsMetadataParam struct {
	BaseParam
	Params RefreshGuestOsMetadataDetailParam `json:"params"`
}
