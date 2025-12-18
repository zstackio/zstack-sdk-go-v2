// Copyright (c) ZStack.io, Inc.

package param

// AddExternalPrimaryStorageDetailParam AddExternalPrimaryStorage详细参数
type AddExternalPrimaryStorageDetailParam struct {
	rest string `json:"identity" validate:"required"` // 必填
	rest string `json:"defaultOutputProtocol" validate:"required"` // 必填
	rest string `json:"config,omitempty"`
	rest string `json:"url" validate:"required"` // 必填
	rest string `json:"name" validate:"required"` // 必填
	rest string `json:"description,omitempty"`
	rest string `json:"type,omitempty"`
	rest string `json:"zoneUuid" validate:"required"` // 必填
	rest string `json:"resourceUuid,omitempty"`
	rest []string `json:"tagUuids,omitempty"`
}

// AddExternalPrimaryStorageParam AddExternalPrimaryStorage请求参数
type AddExternalPrimaryStorageParam struct {
	BaseParam
	Params AddExternalPrimaryStorageDetailParam `json:"params"` // 详细参数
}

