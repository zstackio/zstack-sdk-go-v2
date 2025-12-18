// Copyright (c) ZStack.io, Inc.

package param

// CreateVmInstanceFromOvfDetailParam CreateVmInstanceFromOvf详细参数
type CreateVmInstanceFromOvfDetailParam struct {
	rest string `json:"xmlBase64" validate:"required"` // 必填
	rest string `json:"jsonImageInfos" validate:"required"` // 必填
	rest string `json:"backupStorageUuid" validate:"required"` // 必填
	rest string `json:"jsonCreateVmParam" validate:"required"` // 必填
	rest bool `json:"deleteImageAfterSuccess,omitempty"`
	rest bool `json:"deleteImageOnFail,omitempty"`
	rest string `json:"resourceUuid,omitempty"`
	rest []string `json:"tagUuids,omitempty"`
}

// CreateVmInstanceFromOvfParam CreateVmInstanceFromOvf请求参数
type CreateVmInstanceFromOvfParam struct {
	BaseParam
	Params CreateVmInstanceFromOvfDetailParam `json:"params"` // 详细参数
}

