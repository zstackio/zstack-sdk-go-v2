// Copyright (c) ZStack.io, Inc.

package param

// CreateBuildAppDetailParam CreateBuildApp detail param
type CreateBuildAppDetailParam struct {
	BuildSystemUuid string `json:"buildSystemUuid" validate:"required"`
	BackupStorageUuid string `json:"backupStorageUuid" validate:"required"`
	DataPath string `json:"dataPath" validate:"required"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateBuildAppParam CreateBuildApp request param
type CreateBuildAppParam struct {
	BaseParam
	Params CreateBuildAppDetailParam `json:"params"`
}
