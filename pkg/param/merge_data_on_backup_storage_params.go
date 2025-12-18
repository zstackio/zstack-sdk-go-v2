// Copyright (c) ZStack.io, Inc.

package param

// MergeDataOnBackupStorageDetailParam MergeDataOnBackupStorage详细参数
type MergeDataOnBackupStorageDetailParam struct {
	rest string `json:"backupStorageUuid" validate:"required"` // 必填
}

// MergeDataOnBackupStorageParam MergeDataOnBackupStorage请求参数
type MergeDataOnBackupStorageParam struct {
	BaseParam
	Params MergeDataOnBackupStorageDetailParam `json:"params"` // 详细参数
}

