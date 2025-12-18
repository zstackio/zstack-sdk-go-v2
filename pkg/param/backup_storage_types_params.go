// Copyright (c) ZStack.io, Inc.

package param

// GetBackupStorageTypesDetailParam GetBackupStorageTypes详细参数
type GetBackupStorageTypesDetailParam struct {
}

// GetBackupStorageTypesParam GetBackupStorageTypes请求参数
type GetBackupStorageTypesParam struct {
	BaseParam
	Params GetBackupStorageTypesDetailParam `json:"params"` // 详细参数
}

