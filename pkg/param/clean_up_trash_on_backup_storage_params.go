// Copyright (c) ZStack.io, Inc.

package param

// CleanUpTrashOnBackupStorageDetailParam CleanUpTrashOnBackupStorage detail param
type CleanUpTrashOnBackupStorageDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	TrashId int64 `json:"trashId,omitempty"`
}

// CleanUpTrashOnBackupStorageParam CleanUpTrashOnBackupStorage request param
type CleanUpTrashOnBackupStorageParam struct {
	BaseParam
	Params CleanUpTrashOnBackupStorageDetailParam `json:"params"`
}
