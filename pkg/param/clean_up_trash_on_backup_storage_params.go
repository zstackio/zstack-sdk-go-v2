// Copyright (c) ZStack.io, Inc.

package param

// CleanUpTrashOnBackupStorageDetailParam CleanUpTrashOnBackupStorage详细参数
type CleanUpTrashOnBackupStorageDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest int64 `json:"trashId,omitempty"`
}

// CleanUpTrashOnBackupStorageParam CleanUpTrashOnBackupStorage请求参数
type CleanUpTrashOnBackupStorageParam struct {
	BaseParam
	Params CleanUpTrashOnBackupStorageDetailParam `json:"params"` // 详细参数
}

