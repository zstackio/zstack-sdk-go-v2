// Copyright (c) ZStack.io, Inc.

package param

// GetInterdependentL3NetworksBackupStoragesDetailParam GetInterdependentL3NetworksBackupStorages detail param
type GetInterdependentL3NetworksBackupStoragesDetailParam struct {
	ZoneUuid string `json:"zoneUuid" validate:"required"`
	BackupStorageUuid string `json:"backupStorageUuid,omitempty"`
	L3NetworkUuids []string `json:"l3NetworkUuids,omitempty"`
}

// GetInterdependentL3NetworksBackupStoragesParam GetInterdependentL3NetworksBackupStorages request param
type GetInterdependentL3NetworksBackupStoragesParam struct {
	BaseParam
	Params GetInterdependentL3NetworksBackupStoragesDetailParam `json:"params"`
}
