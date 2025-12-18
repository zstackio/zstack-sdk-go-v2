// Copyright (c) ZStack.io, Inc.

package param

// GetPrimaryStorageCandidatesForVmMigrationDetailParam GetPrimaryStorageCandidatesForVmMigration detail param
type GetPrimaryStorageCandidatesForVmMigrationDetailParam struct {
	VmInstanceUuid string `json:"vmInstanceUuid" validate:"required"`
	WithDataVolumes bool `json:"withDataVolumes,omitempty"`
	MigrateStorageOnly bool `json:"migrateStorageOnly,omitempty"`
}

// GetPrimaryStorageCandidatesForVmMigrationParam GetPrimaryStorageCandidatesForVmMigration request param
type GetPrimaryStorageCandidatesForVmMigrationParam struct {
	BaseParam
	Params GetPrimaryStorageCandidatesForVmMigrationDetailParam `json:"params"`
}
