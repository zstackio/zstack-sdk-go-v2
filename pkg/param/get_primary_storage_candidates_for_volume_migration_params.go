// Copyright (c) ZStack.io, Inc.

package param

// GetPrimaryStorageCandidatesForVolumeMigrationDetailParam GetPrimaryStorageCandidatesForVolumeMigration detail param
type GetPrimaryStorageCandidatesForVolumeMigrationDetailParam struct {
	VolumeUuid string `json:"volumeUuid" validate:"required"`
}

// GetPrimaryStorageCandidatesForVolumeMigrationParam GetPrimaryStorageCandidatesForVolumeMigration request param
type GetPrimaryStorageCandidatesForVolumeMigrationParam struct {
	BaseParam
	Params GetPrimaryStorageCandidatesForVolumeMigrationDetailParam `json:"params"`
}
