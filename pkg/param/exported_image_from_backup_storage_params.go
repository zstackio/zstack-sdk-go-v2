// Copyright (c) ZStack.io, Inc.

package param

// DeleteExportedImageFromBackupStorageDetailParam DeleteExportedImageFromBackupStorage详细参数
type DeleteExportedImageFromBackupStorageDetailParam struct {
	rest string `json:"backupStorageUuid" validate:"required"` // 必填
	rest string `json:"imageUuid" validate:"required"` // 必填
}

// DeleteExportedImageFromBackupStorageParam DeleteExportedImageFromBackupStorage请求参数
type DeleteExportedImageFromBackupStorageParam struct {
	BaseParam
	Params DeleteExportedImageFromBackupStorageDetailParam `json:"params"` // 详细参数
}

