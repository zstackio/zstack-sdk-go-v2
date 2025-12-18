// Copyright (c) ZStack.io, Inc.

package param

// GetPrimaryStorageCandidatesForVolumeMigrationDetailParam GetPrimaryStorageCandidatesForVolumeMigration详细参数
type GetPrimaryStorageCandidatesForVolumeMigrationDetailParam struct {
	rest string `json:"volumeUuid" validate:"required"` // 必填
}

// GetPrimaryStorageCandidatesForVolumeMigrationParam GetPrimaryStorageCandidatesForVolumeMigration请求参数
type GetPrimaryStorageCandidatesForVolumeMigrationParam struct {
	BaseParam
	Params GetPrimaryStorageCandidatesForVolumeMigrationDetailParam `json:"params"` // 详细参数
}

