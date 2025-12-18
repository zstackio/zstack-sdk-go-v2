// Copyright (c) ZStack.io, Inc.

package param

// CreateEcsImageFromLocalImageDetailParam CreateEcsImageFromLocalImage detail param
type CreateEcsImageFromLocalImageDetailParam struct {
	ImageUuid string `json:"imageUuid" validate:"required"`
	DataCenterUuid string `json:"dataCenterUuid" validate:"required"`
	BackupStorageUuid string `json:"backupStorageUuid,omitempty"`
	Description string `json:"description,omitempty"`
	Name string `json:"name" validate:"required"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateEcsImageFromLocalImageParam CreateEcsImageFromLocalImage request param
type CreateEcsImageFromLocalImageParam struct {
	BaseParam
	Params CreateEcsImageFromLocalImageDetailParam `json:"params"`
}
