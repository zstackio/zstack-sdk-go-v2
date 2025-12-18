// Copyright (c) ZStack.io, Inc.

package param

// ExportVmOvaPackageDetailParam ExportVmOvaPackage详细参数
type ExportVmOvaPackageDetailParam struct {
	rest string `json:"name,omitempty"`
	rest string `json:"description,omitempty"`
	rest string `json:"vmUuid" validate:"required"` // 必填
	rest string `json:"backupStorageUuid" validate:"required"` // 必填
	rest string `json:"resourceUuid,omitempty"`
	rest []string `json:"tagUuids,omitempty"`
}

// ExportVmOvaPackageParam ExportVmOvaPackage请求参数
type ExportVmOvaPackageParam struct {
	BaseParam
	Params ExportVmOvaPackageDetailParam `json:"params"` // 详细参数
}

