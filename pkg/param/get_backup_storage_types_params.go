// Copyright (c) ZStack.io, Inc.

package param

// GetBackupStorageTypesDetailParam GetBackupStorageTypes detail param
type GetBackupStorageTypesDetailParam struct {
}

// GetBackupStorageTypesParam GetBackupStorageTypes request param
type GetBackupStorageTypesParam struct {
	BaseParam
	Params GetBackupStorageTypesDetailParam `json:"params"`
}
