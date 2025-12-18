// Copyright (c) ZStack.io, Inc.

package param

// GetTrashOnPrimaryStorageDetailParam GetTrashOnPrimaryStorage详细参数
type GetTrashOnPrimaryStorageDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"resourceUuid,omitempty"`
	rest string `json:"resourceType,omitempty"`
	rest string `json:"trashType,omitempty"`
}

// GetTrashOnPrimaryStorageParam GetTrashOnPrimaryStorage请求参数
type GetTrashOnPrimaryStorageParam struct {
	BaseParam
	Params GetTrashOnPrimaryStorageDetailParam `json:"params"` // 详细参数
}

