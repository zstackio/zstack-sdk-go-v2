// Copyright (c) ZStack.io, Inc.

package param

// AddModelServiceDetailParam AddModelService detail param
type AddModelServiceDetailParam struct {
	Name string `json:"name" validate:"required"`
	Description string `json:"description,omitempty"`
	Yaml string `json:"yaml" validate:"required"`
	ModelCenterUuid string `json:"modelCenterUuid" validate:"required"`
	RequestCpu int `json:"requestCpu" validate:"required"`
	RequestMemory int64 `json:"requestMemory" validate:"required"`
	GpuComputeCapability string `json:"gpuComputeCapability,omitempty"`
	InstallPath string `json:"installPath" validate:"required"`
	System bool `json:"system,omitempty"`
	StartCommand string `json:"startCommand" validate:"required"`
	ContainerCommand string `json:"containerCommand,omitempty"`
	ContainerArgs string `json:"containerArgs,omitempty"`
	PythonVersion string `json:"pythonVersion,omitempty"`
	CondaVersion string `json:"condaVersion,omitempty"`
	ZoneUuid string `json:"zoneUuid,omitempty"`
	Type string `json:"type,omitempty"`
	Source string `json:"source,omitempty"`
	Framework string `json:"framework,omitempty"`
	ModelUuids []string `json:"modelUuids,omitempty"`
	ArchitectureImages []interface{} `json:"architectureImages,omitempty"`
	SupportDistributed bool `json:"supportDistributed,omitempty"`
	VendorToSpecUuidsMap map[string]interface{} `json:"vendorToSpecUuidsMap,omitempty"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// AddModelServiceParam AddModelService request param
type AddModelServiceParam struct {
	BaseParam
	Params AddModelServiceDetailParam `json:"params"`
}
