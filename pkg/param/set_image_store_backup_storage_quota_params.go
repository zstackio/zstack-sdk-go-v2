// Copyright (c) ZStack.io, Inc.

package param

// SetImageStoreBackupStorageQuotaDetailParam SetImageStoreBackupStorageQuota详细参数
type SetImageStoreBackupStorageQuotaDetailParam struct {
	rest []string `json:"uuids,omitempty"`
	rest int64 `json:"maxCapacity" validate:"required"` // 必填
}

// SetImageStoreBackupStorageQuotaParam SetImageStoreBackupStorageQuota请求参数
type SetImageStoreBackupStorageQuotaParam struct {
	BaseParam
	Params SetImageStoreBackupStorageQuotaDetailParam `json:"params"` // 详细参数
}

