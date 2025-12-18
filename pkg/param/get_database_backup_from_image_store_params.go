// Copyright (c) ZStack.io, Inc.

package param

// GetDatabaseBackupFromImageStoreDetailParam GetDatabaseBackupFromImageStore detail param
type GetDatabaseBackupFromImageStoreDetailParam struct {
	Url string `json:"url" validate:"required"`
	RegistryPort int `json:"registryPort,omitempty"`
}

// GetDatabaseBackupFromImageStoreParam GetDatabaseBackupFromImageStore request param
type GetDatabaseBackupFromImageStoreParam struct {
	BaseParam
	Params GetDatabaseBackupFromImageStoreDetailParam `json:"params"`
}
