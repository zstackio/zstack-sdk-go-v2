// Copyright (c) ZStack.io, Inc.

package param

// GetTrashOnBackupStorageDetailParam GetTrashOnBackupStorage详细参数
type GetTrashOnBackupStorageDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"resourceUuid,omitempty"`
	rest string `json:"resourceType,omitempty"`
	rest string `json:"trashType,omitempty"`
}

// GetTrashOnBackupStorageParam GetTrashOnBackupStorage请求参数
type GetTrashOnBackupStorageParam struct {
	BaseParam
	Params GetTrashOnBackupStorageDetailParam `json:"params"` // 详细参数
}

