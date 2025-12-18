// Copyright (c) ZStack.io, Inc.

package param

// AddBlockPrimaryStorageDetailParam AddBlockPrimaryStorage详细参数
type AddBlockPrimaryStorageDetailParam struct {
	rest string `json:"vendorName" validate:"required"` // 必填
	rest string `json:"metadata" validate:"required"` // 必填
	rest string `json:"url" validate:"required"` // 必填
	rest string `json:"name" validate:"required"` // 必填
	rest string `json:"description,omitempty"`
	rest string `json:"type,omitempty"`
	rest string `json:"zoneUuid" validate:"required"` // 必填
	rest string `json:"resourceUuid,omitempty"`
	rest []string `json:"tagUuids,omitempty"`
}

// AddBlockPrimaryStorageParam AddBlockPrimaryStorage请求参数
type AddBlockPrimaryStorageParam struct {
	BaseParam
	Params AddBlockPrimaryStorageDetailParam `json:"params"` // 详细参数
}

