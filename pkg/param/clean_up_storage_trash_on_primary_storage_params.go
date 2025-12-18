// Copyright (c) ZStack.io, Inc.

package param

// CleanUpStorageTrashOnPrimaryStorageDetailParam CleanUpStorageTrashOnPrimaryStorage详细参数
type CleanUpStorageTrashOnPrimaryStorageDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest bool `json:"force,omitempty"`
}

// CleanUpStorageTrashOnPrimaryStorageParam CleanUpStorageTrashOnPrimaryStorage请求参数
type CleanUpStorageTrashOnPrimaryStorageParam struct {
	BaseParam
	Params CleanUpStorageTrashOnPrimaryStorageDetailParam `json:"params"` // 详细参数
}

