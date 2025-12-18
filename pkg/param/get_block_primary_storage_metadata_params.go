// Copyright (c) ZStack.io, Inc.

package param

// GetBlockPrimaryStorageMetadataDetailParam GetBlockPrimaryStorageMetadata detail param
type GetBlockPrimaryStorageMetadataDetailParam struct {
	VendorName string `json:"vendorName" validate:"required"`
	Metadata string `json:"metadata" validate:"required"`
}

// GetBlockPrimaryStorageMetadataParam GetBlockPrimaryStorageMetadata request param
type GetBlockPrimaryStorageMetadataParam struct {
	BaseParam
	Params GetBlockPrimaryStorageMetadataDetailParam `json:"params"`
}
