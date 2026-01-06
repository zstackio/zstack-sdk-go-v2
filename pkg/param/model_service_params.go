// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now // avoid unused import

// CloneModelServiceParamDetail CloneModelService detail param
type CloneModelServiceParamDetail struct {
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
	ArchitectureImages []ArchitectureImageMappingParam `json:"architectureImages,omitempty"`
	SupportDistributed bool `json:"supportDistributed,omitempty"`
	VendorToSpecUuidsMap map[string]interface{} `json:"vendorToSpecUuidsMap,omitempty"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CloneModelServiceParam CloneModelService request param
type CloneModelServiceParam struct {
	BaseParam
	Params CloneModelServiceParamDetail `json:"params"`
}
// UpdateModelServiceParamDetail UpdateModelService detail param
type UpdateModelServiceParamDetail struct {
	Uuid string `json:"uuid" validate:"required"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Yaml string `json:"yaml,omitempty"`
	RequestCpu int `json:"requestCpu,omitempty"`
	RequestMemory int64 `json:"requestMemory,omitempty"`
	GpuComputeCapability string `json:"gpuComputeCapability,omitempty"`
	StartCommand string `json:"startCommand,omitempty"`
	ContainerCommand string `json:"containerCommand,omitempty"`
	ContainerArgs string `json:"containerArgs,omitempty"`
	PythonVersion string `json:"pythonVersion,omitempty"`
	Type string `json:"type,omitempty"`
	Source string `json:"source,omitempty"`
	Framework string `json:"framework,omitempty"`
	ArchitectureImages []ArchitectureImageMappingParam `json:"architectureImages,omitempty"`
	SupportDistributed bool `json:"supportDistributed,omitempty"`
	EnvironmentParameters map[string]string `json:"environmentParameters,omitempty"`
	StartupParameters map[string]string `json:"startupParameters,omitempty"`
	InferenceParams map[string]string `json:"inferenceParams,omitempty"`
	ServiceName string `json:"serviceName,omitempty"`
	ServicePorts []int `json:"servicePorts,omitempty"`
	ServiceLivez string `json:"serviceLivez,omitempty"`
	ServiceReadyz string `json:"serviceReadyz,omitempty"`
	ServiceBootupTime int `json:"serviceBootupTime,omitempty"`
	ServiceInstallPath string `json:"serviceInstallPath,omitempty"`
	ServiceStartCommand string `json:"serviceStartCommand,omitempty"`
	VendorToSpecUuidsMap map[string]interface{} `json:"vendorToSpecUuidsMap,omitempty"`
}

// UpdateModelServiceParam UpdateModelService request param
type UpdateModelServiceParam struct {
	BaseParam
	Params UpdateModelServiceParamDetail `json:"params"`
}
// DeleteModelServiceParamDetail DeleteModelService detail param
type DeleteModelServiceParamDetail struct {
	Uuid string `json:"uuid" validate:"required"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// DeleteModelServiceParam DeleteModelService request param
type DeleteModelServiceParam struct {
	BaseParam
	Params DeleteModelServiceParamDetail `json:"params"`
}
// AddModelServiceParamDetail AddModelService detail param
type AddModelServiceParamDetail struct {
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
	ArchitectureImages []ArchitectureImageMappingParam `json:"architectureImages,omitempty"`
	SupportDistributed bool `json:"supportDistributed,omitempty"`
	VendorToSpecUuidsMap map[string]interface{} `json:"vendorToSpecUuidsMap,omitempty"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// AddModelServiceParam AddModelService request param
type AddModelServiceParam struct {
	BaseParam
	Params AddModelServiceParamDetail `json:"params"`
}
