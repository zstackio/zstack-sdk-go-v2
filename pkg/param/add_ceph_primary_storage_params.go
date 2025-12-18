// Copyright (c) ZStack.io, Inc.

package param

// AddCephPrimaryStorageDetailParam AddCephPrimaryStorage详细参数
type AddCephPrimaryStorageDetailParam struct {
	rest []string `json:"monUrls" validate:"required"` // 必填
	rest string `json:"rootVolumePoolName,omitempty"`
	rest string `json:"dataVolumePoolName,omitempty"`
	rest string `json:"imageCachePoolName,omitempty"`
	rest string `json:"url,omitempty"` // 必填
	rest string `json:"name" validate:"required"` // 必填
	rest string `json:"description,omitempty"`
	rest string `json:"type,omitempty"`
	rest string `json:"zoneUuid" validate:"required"` // 必填
	rest string `json:"resourceUuid,omitempty"`
	rest []string `json:"tagUuids,omitempty"`
}

// AddCephPrimaryStorageParam AddCephPrimaryStorage请求参数
type AddCephPrimaryStorageParam struct {
	BaseParam
	Params AddCephPrimaryStorageDetailParam `json:"params"` // 详细参数
}

