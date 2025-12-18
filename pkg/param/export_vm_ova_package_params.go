// Copyright (c) ZStack.io, Inc.

package param

// ExportVmOvaPackageDetailParam ExportVmOvaPackage detail param
type ExportVmOvaPackageDetailParam struct {
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	VmUuid string `json:"vmUuid" validate:"required"`
	BackupStorageUuid string `json:"backupStorageUuid" validate:"required"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// ExportVmOvaPackageParam ExportVmOvaPackage request param
type ExportVmOvaPackageParam struct {
	BaseParam
	Params ExportVmOvaPackageDetailParam `json:"params"`
}
