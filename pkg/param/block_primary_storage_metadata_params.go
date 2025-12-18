// Copyright (c) ZStack.io, Inc.

package param

// GetBlockPrimaryStorageMetadataDetailParam GetBlockPrimaryStorageMetadata详细参数
type GetBlockPrimaryStorageMetadataDetailParam struct {
	rest string `json:"vendorName" validate:"required"` // 必填
	rest string `json:"metadata" validate:"required"` // 必填
}

// GetBlockPrimaryStorageMetadataParam GetBlockPrimaryStorageMetadata请求参数
type GetBlockPrimaryStorageMetadataParam struct {
	BaseParam
	Params GetBlockPrimaryStorageMetadataDetailParam `json:"params"` // 详细参数
}

