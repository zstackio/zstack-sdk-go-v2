// Copyright (c) ZStack.io, Inc.

package param

// AddSharedBlockGroupPrimaryStorageDetailParam AddSharedBlockGroupPrimaryStorage详细参数
type AddSharedBlockGroupPrimaryStorageDetailParam struct {
	rest []string `json:"diskUuids" validate:"required"` // 必填
	rest string `json:"url,omitempty"` // 必填
	rest string `json:"name" validate:"required"` // 必填
	rest string `json:"description,omitempty"`
	rest string `json:"type,omitempty"`
	rest string `json:"zoneUuid" validate:"required"` // 必填
	rest string `json:"resourceUuid,omitempty"`
	rest []string `json:"tagUuids,omitempty"`
}

// AddSharedBlockGroupPrimaryStorageParam AddSharedBlockGroupPrimaryStorage请求参数
type AddSharedBlockGroupPrimaryStorageParam struct {
	BaseParam
	Params AddSharedBlockGroupPrimaryStorageDetailParam `json:"params"` // 详细参数
}

