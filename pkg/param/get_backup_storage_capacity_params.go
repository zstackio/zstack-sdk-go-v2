// Copyright (c) ZStack.io, Inc.

package param

// GetBackupStorageCapacityDetailParam GetBackupStorageCapacity detail param
type GetBackupStorageCapacityDetailParam struct {
	ZoneUuids []string `json:"zoneUuids,omitempty"`
	BackupStorageUuids []string `json:"backupStorageUuids,omitempty"`
	All bool `json:"all,omitempty"`
}

// GetBackupStorageCapacityParam GetBackupStorageCapacity request param
type GetBackupStorageCapacityParam struct {
	BaseParam
	Params GetBackupStorageCapacityDetailParam `json:"params"`
}
