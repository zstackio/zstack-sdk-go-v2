// Copyright (c) ZStack.io, Inc.

package param

// GetGuestOsMetadataDetailParam GetGuestOsMetadata详细参数
type GetGuestOsMetadataDetailParam struct {
}

// GetGuestOsMetadataParam GetGuestOsMetadata请求参数
type GetGuestOsMetadataParam struct {
	BaseParam
	Params GetGuestOsMetadataDetailParam `json:"params"` // 详细参数
}

