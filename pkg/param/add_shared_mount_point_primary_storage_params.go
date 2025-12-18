// Copyright (c) ZStack.io, Inc.

package param

// AddSharedMountPointPrimaryStorageDetailParam AddSharedMountPointPrimaryStorage详细参数
type AddSharedMountPointPrimaryStorageDetailParam struct {
	rest string `json:"url" validate:"required"` // 必填
	rest string `json:"name" validate:"required"` // 必填
	rest string `json:"description,omitempty"`
	rest string `json:"type,omitempty"`
	rest string `json:"zoneUuid" validate:"required"` // 必填
	rest string `json:"resourceUuid,omitempty"`
	rest []string `json:"tagUuids,omitempty"`
}

// AddSharedMountPointPrimaryStorageParam AddSharedMountPointPrimaryStorage请求参数
type AddSharedMountPointPrimaryStorageParam struct {
	BaseParam
	Params AddSharedMountPointPrimaryStorageDetailParam `json:"params"` // 详细参数
}

