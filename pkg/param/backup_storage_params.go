// Copyright (c) ZStack.io, Inc.

package param

// DeleteBackupStorageDetailParam DeleteBackupStorage详细参数
type DeleteBackupStorageDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"deleteMode,omitempty"`
}

// DeleteBackupStorageParam DeleteBackupStorage请求参数
type DeleteBackupStorageParam struct {
	BaseParam
	Params DeleteBackupStorageDetailParam `json:"params"` // 详细参数
}

