// Copyright (c) ZStack.io, Inc.

package param

// ExportImageFromBackupStorageDetailParam ExportImageFromBackupStorage详细参数
type ExportImageFromBackupStorageDetailParam struct {
	rest string `json:"backupStorageUuid" validate:"required"` // 必填
	rest string `json:"imageUuid" validate:"required"` // 必填
	rest string `json:"exportFormat,omitempty"`
}

// ExportImageFromBackupStorageParam ExportImageFromBackupStorage请求参数
type ExportImageFromBackupStorageParam struct {
	BaseParam
	Params ExportImageFromBackupStorageDetailParam `json:"params"` // 详细参数
}

