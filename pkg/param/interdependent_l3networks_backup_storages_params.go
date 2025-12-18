// Copyright (c) ZStack.io, Inc.

package param

// GetInterdependentL3NetworksBackupStoragesDetailParam GetInterdependentL3NetworksBackupStorages详细参数
type GetInterdependentL3NetworksBackupStoragesDetailParam struct {
	rest string `json:"zoneUuid" validate:"required"` // 必填
	rest string `json:"backupStorageUuid,omitempty"`
	rest []string `json:"l3NetworkUuids,omitempty"`
}

// GetInterdependentL3NetworksBackupStoragesParam GetInterdependentL3NetworksBackupStorages请求参数
type GetInterdependentL3NetworksBackupStoragesParam struct {
	BaseParam
	Params GetInterdependentL3NetworksBackupStoragesDetailParam `json:"params"` // 详细参数
}

