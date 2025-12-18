// Copyright (c) ZStack.io, Inc.

package param

// AddLocalPrimaryStorageDetailParam AddLocalPrimaryStorage详细参数
type AddLocalPrimaryStorageDetailParam struct {
	rest string `json:"url" validate:"required"` // 必填
	rest string `json:"name" validate:"required"` // 必填
	rest string `json:"description,omitempty"`
	rest string `json:"type,omitempty"`
	rest string `json:"zoneUuid" validate:"required"` // 必填
	rest string `json:"resourceUuid,omitempty"`
	rest []string `json:"tagUuids,omitempty"`
}

// AddLocalPrimaryStorageParam AddLocalPrimaryStorage请求参数
type AddLocalPrimaryStorageParam struct {
	BaseParam
	Params AddLocalPrimaryStorageDetailParam `json:"params"` // 详细参数
}

