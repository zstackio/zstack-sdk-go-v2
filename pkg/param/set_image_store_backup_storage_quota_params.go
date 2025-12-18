// Copyright (c) ZStack.io, Inc.

package param

// SetImageStoreBackupStorageQuotaDetailParam SetImageStoreBackupStorageQuota detail param
type SetImageStoreBackupStorageQuotaDetailParam struct {
	Uuids []string `json:"uuids,omitempty"`
	MaxCapacity int64 `json:"maxCapacity" validate:"required"`
}

// SetImageStoreBackupStorageQuotaParam SetImageStoreBackupStorageQuota request param
type SetImageStoreBackupStorageQuotaParam struct {
	BaseParam
	Params SetImageStoreBackupStorageQuotaDetailParam `json:"params"`
}
