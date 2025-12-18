// Copyright (c) ZStack.io, Inc.

package param

// ReconnectImageStoreBackupStorageDetailParam ReconnectImageStoreBackupStorage详细参数
type ReconnectImageStoreBackupStorageDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
}

// ReconnectImageStoreBackupStorageParam ReconnectImageStoreBackupStorage请求参数
type ReconnectImageStoreBackupStorageParam struct {
	BaseParam
	Params ReconnectImageStoreBackupStorageDetailParam `json:"params"` // 详细参数
}

