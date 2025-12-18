// Copyright (c) ZStack.io, Inc.

package param

// UpdateBlockPrimaryStorageDetailParam UpdateBlockPrimaryStorage详细参数
type UpdateBlockPrimaryStorageDetailParam struct {
	rest string `json:"vendorName,omitempty"`
	rest string `json:"metadata,omitempty"`
	rest string `json:"uuid" validate:"required"` // 必填
}

// UpdateBlockPrimaryStorageParam UpdateBlockPrimaryStorage请求参数
type UpdateBlockPrimaryStorageParam struct {
	BaseParam
	Params UpdateBlockPrimaryStorageDetailParam `json:"params"` // 详细参数
}

