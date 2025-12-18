// Copyright (c) ZStack.io, Inc.

package param

// ReconnectBackupStorageDetailParam ReconnectBackupStorage详细参数
type ReconnectBackupStorageDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
}

// ReconnectBackupStorageParam ReconnectBackupStorage请求参数
type ReconnectBackupStorageParam struct {
	BaseParam
	Params ReconnectBackupStorageDetailParam `json:"params"` // 详细参数
}

