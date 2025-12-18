// Copyright (c) ZStack.io, Inc.

package param

// GetHostCandidatesForVmMigrationDetailParam GetHostCandidatesForVmMigration detail param
type GetHostCandidatesForVmMigrationDetailParam struct {
	VmInstanceUuid string `json:"vmInstanceUuid" validate:"required"`
	DstPrimaryStorageUuid string `json:"dstPrimaryStorageUuid" validate:"required"`
	Limit int `json:"limit,omitempty"`
}

// GetHostCandidatesForVmMigrationParam GetHostCandidatesForVmMigration request param
type GetHostCandidatesForVmMigrationParam struct {
	BaseParam
	Params GetHostCandidatesForVmMigrationDetailParam `json:"params"`
}
