// Copyright (c) ZStack.io, Inc.

package param

// RefreshGuestOsMetadataDetailParam RefreshGuestOsMetadata详细参数
type RefreshGuestOsMetadataDetailParam struct {
}

// RefreshGuestOsMetadataParam RefreshGuestOsMetadata请求参数
type RefreshGuestOsMetadataParam struct {
	BaseParam
	Params RefreshGuestOsMetadataDetailParam `json:"params"` // 详细参数
}

