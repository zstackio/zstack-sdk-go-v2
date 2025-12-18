// Copyright (c) ZStack.io, Inc.

package param

// GetHostCandidatesForVmMigrationDetailParam GetHostCandidatesForVmMigration详细参数
type GetHostCandidatesForVmMigrationDetailParam struct {
	rest string `json:"vmInstanceUuid" validate:"required"` // 必填
	rest string `json:"dstPrimaryStorageUuid" validate:"required"` // 必填
	rest int `json:"limit,omitempty"`
}

// GetHostCandidatesForVmMigrationParam GetHostCandidatesForVmMigration请求参数
type GetHostCandidatesForVmMigrationParam struct {
	BaseParam
	Params GetHostCandidatesForVmMigrationDetailParam `json:"params"` // 详细参数
}

