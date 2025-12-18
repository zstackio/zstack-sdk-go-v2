// Copyright (c) ZStack.io, Inc.

package param

// GetPrimaryStorageCandidatesForVmMigrationDetailParam GetPrimaryStorageCandidatesForVmMigration详细参数
type GetPrimaryStorageCandidatesForVmMigrationDetailParam struct {
	rest string `json:"vmInstanceUuid" validate:"required"` // 必填
	rest bool `json:"withDataVolumes,omitempty"`
	rest bool `json:"migrateStorageOnly,omitempty"`
}

// GetPrimaryStorageCandidatesForVmMigrationParam GetPrimaryStorageCandidatesForVmMigration请求参数
type GetPrimaryStorageCandidatesForVmMigrationParam struct {
	BaseParam
	Params GetPrimaryStorageCandidatesForVmMigrationDetailParam `json:"params"` // 详细参数
}

