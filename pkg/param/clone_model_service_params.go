// Copyright (c) ZStack.io, Inc.

package param

// CloneModelServiceDetailParam CloneModelService detail param
type CloneModelServiceDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Readme string `json:"readme,omitempty"`
	Size int64 `json:"size,omitempty"`
	GpuComputeCapability string `json:"gpuComputeCapability,omitempty"`
	InstallPath string `json:"installPath,omitempty"`
	StartCommand string `json:"startCommand,omitempty"`
	ContainerCommand string `json:"containerCommand,omitempty"`
	ContainerArgs string `json:"containerArgs,omitempty"`
	PythonVersion string `json:"pythonVersion,omitempty"`
	CondaVersion string `json:"condaVersion,omitempty"`
	System bool `json:"system,omitempty"`
	Type string `json:"type,omitempty"`
	Yaml string `json:"yaml,omitempty"`
	Source string `json:"source,omitempty"`
	Framework string `json:"framework,omitempty"`
	RequestCpu int `json:"requestCpu,omitempty"`
	RequestMemory int64 `json:"requestMemory,omitempty"`
	CpuArchitectures []string `json:"cpuArchitectures,omitempty"`
	ArchitectureImages []interface{} `json:"architectureImages,omitempty"`
	SupportDistributed bool `json:"supportDistributed,omitempty"`
	VendorToSpecUuidsMap map[string]interface{} `json:"vendorToSpecUuidsMap,omitempty"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CloneModelServiceParam CloneModelService request param
type CloneModelServiceParam struct {
	BaseParam
	Params CloneModelServiceDetailParam `json:"params"`
}
