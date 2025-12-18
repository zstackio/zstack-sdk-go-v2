// Copyright (c) ZStack.io, Inc.

package param

// UploadFileToVmDetailParam UploadFileToVm detail param
type UploadFileToVmDetailParam struct {
	VmInstanceUuids []string `json:"vmInstanceUuids" validate:"required"`
	FileContent string `json:"fileContent" validate:"required"`
	RemotePath string `json:"remotePath" validate:"required"`
}

// UploadFileToVmParam UploadFileToVm request param
type UploadFileToVmParam struct {
	BaseParam
	Params UploadFileToVmDetailParam `json:"params"`
}
