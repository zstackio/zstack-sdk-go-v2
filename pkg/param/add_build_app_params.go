// Copyright (c) ZStack.io, Inc.

package param

// AddBuildAppDetailParam AddBuildApp detail param
type AddBuildAppDetailParam struct {
	Url string `json:"url" validate:"required"`
	Type string `json:"type,omitempty"`
	BackupStorageUuid string `json:"backupStorageUuid,omitempty"`
	Hostname string `json:"hostname,omitempty"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// AddBuildAppParam AddBuildApp request param
type AddBuildAppParam struct {
	BaseParam
	Params AddBuildAppDetailParam `json:"params"`
}
