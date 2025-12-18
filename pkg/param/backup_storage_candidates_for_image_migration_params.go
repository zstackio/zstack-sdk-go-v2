// Copyright (c) ZStack.io, Inc.

package param

// GetBackupStorageCandidatesForImageMigrationDetailParam GetBackupStorageCandidatesForImageMigration详细参数
type GetBackupStorageCandidatesForImageMigrationDetailParam struct {
	rest string `json:"srcBackupStorageUuid" validate:"required"` // 必填
}

// GetBackupStorageCandidatesForImageMigrationParam GetBackupStorageCandidatesForImageMigration请求参数
type GetBackupStorageCandidatesForImageMigrationParam struct {
	BaseParam
	Params GetBackupStorageCandidatesForImageMigrationDetailParam `json:"params"` // 详细参数
}

