// Copyright (c) ZStack.io, Inc.

package param

// GetBackupStorageCandidatesForImageMigrationDetailParam GetBackupStorageCandidatesForImageMigration detail param
type GetBackupStorageCandidatesForImageMigrationDetailParam struct {
	SrcBackupStorageUuid string `json:"srcBackupStorageUuid" validate:"required"`
}

// GetBackupStorageCandidatesForImageMigrationParam GetBackupStorageCandidatesForImageMigration request param
type GetBackupStorageCandidatesForImageMigrationParam struct {
	BaseParam
	Params GetBackupStorageCandidatesForImageMigrationDetailParam `json:"params"`
}
