// Copyright (c) ZStack.io, Inc.

package param

// AddNfsPrimaryStorageDetailParam AddNfsPrimaryStorage详细参数
type AddNfsPrimaryStorageDetailParam struct {
	rest string `json:"url" validate:"required"` // 必填
	rest string `json:"name" validate:"required"` // 必填
	rest string `json:"description,omitempty"`
	rest string `json:"type,omitempty"`
	rest string `json:"zoneUuid" validate:"required"` // 必填
	rest string `json:"resourceUuid,omitempty"`
	rest []string `json:"tagUuids,omitempty"`
}

// AddNfsPrimaryStorageParam AddNfsPrimaryStorage请求参数
type AddNfsPrimaryStorageParam struct {
	BaseParam
	Params AddNfsPrimaryStorageDetailParam `json:"params"` // 详细参数
}

