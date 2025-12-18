// Copyright (c) ZStack.io, Inc.

package param

// GetBackupStorageCapacityDetailParam GetBackupStorageCapacity详细参数
type GetBackupStorageCapacityDetailParam struct {
	rest []string `json:"zoneUuids,omitempty"`
	rest []string `json:"backupStorageUuids,omitempty"`
	rest bool `json:"all,omitempty"`
}

// GetBackupStorageCapacityParam GetBackupStorageCapacity请求参数
type GetBackupStorageCapacityParam struct {
	BaseParam
	Params GetBackupStorageCapacityDetailParam `json:"params"` // 详细参数
}

