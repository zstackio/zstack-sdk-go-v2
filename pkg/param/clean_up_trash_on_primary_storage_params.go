// Copyright (c) ZStack.io, Inc.

package param

// CleanUpTrashOnPrimaryStorageDetailParam CleanUpTrashOnPrimaryStorage详细参数
type CleanUpTrashOnPrimaryStorageDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest int64 `json:"trashId,omitempty"`
}

// CleanUpTrashOnPrimaryStorageParam CleanUpTrashOnPrimaryStorage请求参数
type CleanUpTrashOnPrimaryStorageParam struct {
	BaseParam
	Params CleanUpTrashOnPrimaryStorageDetailParam `json:"params"` // 详细参数
}

