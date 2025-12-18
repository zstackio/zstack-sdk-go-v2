// Copyright (c) ZStack.io, Inc.

package param

// AddMiniStorageDetailParam AddMiniStorage详细参数
type AddMiniStorageDetailParam struct {
	rest string `json:"diskIdentifier" validate:"required"` // 必填
	rest string `json:"url,omitempty"` // 必填
	rest string `json:"name" validate:"required"` // 必填
	rest string `json:"description,omitempty"`
	rest string `json:"type,omitempty"`
	rest string `json:"zoneUuid" validate:"required"` // 必填
	rest string `json:"resourceUuid,omitempty"`
	rest []string `json:"tagUuids,omitempty"`
}

// AddMiniStorageParam AddMiniStorage请求参数
type AddMiniStorageParam struct {
	BaseParam
	Params AddMiniStorageDetailParam `json:"params"` // 详细参数
}

