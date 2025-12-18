// Copyright (c) ZStack.io, Inc.

package param

// GetCdpBackupStorageRequirementDetailParam GetCdpBackupStorageRequirement detail param
type GetCdpBackupStorageRequirementDetailParam struct {
	BackupStorageUuid string `json:"backupStorageUuid" validate:"required"`
	Limit int `json:"limit,omitempty"`
	Start int `json:"start,omitempty"`
}

// GetCdpBackupStorageRequirementParam GetCdpBackupStorageRequirement request param
type GetCdpBackupStorageRequirementParam struct {
	BaseParam
	Params GetCdpBackupStorageRequirementDetailParam `json:"params"`
}
