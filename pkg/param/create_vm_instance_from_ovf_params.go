// Copyright (c) ZStack.io, Inc.

package param

// CreateVmInstanceFromOvfDetailParam CreateVmInstanceFromOvf detail param
type CreateVmInstanceFromOvfDetailParam struct {
	XmlBase64 string `json:"xmlBase64" validate:"required"`
	JsonImageInfos string `json:"jsonImageInfos" validate:"required"`
	BackupStorageUuid string `json:"backupStorageUuid" validate:"required"`
	JsonCreateVmParam string `json:"jsonCreateVmParam" validate:"required"`
	DeleteImageAfterSuccess bool `json:"deleteImageAfterSuccess,omitempty"`
	DeleteImageOnFail bool `json:"deleteImageOnFail,omitempty"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateVmInstanceFromOvfParam CreateVmInstanceFromOvf request param
type CreateVmInstanceFromOvfParam struct {
	BaseParam
	Params CreateVmInstanceFromOvfDetailParam `json:"params"`
}
