// Copyright (c) ZStack.io, Inc.

package param

// UploadFileToVmDetailParam UploadFileToVm详细参数
type UploadFileToVmDetailParam struct {
	rest []string `json:"vmInstanceUuids" validate:"required"` // 必填
	rest string `json:"fileContent" validate:"required"` // 必填
	rest string `json:"remotePath" validate:"required"` // 必填
}

// UploadFileToVmParam UploadFileToVm请求参数
type UploadFileToVmParam struct {
	BaseParam
	Params UploadFileToVmDetailParam `json:"params"` // 详细参数
}

