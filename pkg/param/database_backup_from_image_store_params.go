// Copyright (c) ZStack.io, Inc.

package param

// GetDatabaseBackupFromImageStoreDetailParam GetDatabaseBackupFromImageStore详细参数
type GetDatabaseBackupFromImageStoreDetailParam struct {
	rest string `json:"url" validate:"required"` // 必填
	rest int `json:"registryPort,omitempty"`
}

// GetDatabaseBackupFromImageStoreParam GetDatabaseBackupFromImageStore请求参数
type GetDatabaseBackupFromImageStoreParam struct {
	BaseParam
	Params GetDatabaseBackupFromImageStoreDetailParam `json:"params"` // 详细参数
}

