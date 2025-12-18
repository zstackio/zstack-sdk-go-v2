// Copyright (c) ZStack.io, Inc.

package param

// CreateZBoxBackupDetailParam CreateZBoxBackup detail param
type CreateZBoxBackupDetailParam struct {
	ZBoxUuid string `json:"zBoxUuid" validate:"required"`
	HostUuids []string `json:"hostUuids,omitempty"`
	BackupStorageUuids []string `json:"backupStorageUuids,omitempty"`
	DryRun bool `json:"dryRun,omitempty"`
	Name string `json:"name" validate:"required"`
	Description string `json:"description,omitempty"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateZBoxBackupParam CreateZBoxBackup request param
type CreateZBoxBackupParam struct {
	BaseParam
	Params CreateZBoxBackupDetailParam `json:"params"`
}
