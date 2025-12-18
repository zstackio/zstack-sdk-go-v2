// Copyright (c) ZStack.io, Inc.

package param

// UpdateModelServiceDetailParam UpdateModelService detail param
type UpdateModelServiceDetailParam struct {
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
	ArchitectureImages []interface{} `json:"architectureImages,omitempty"`
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
	Params UpdateModelServiceDetailParam `json:"params"`
}
