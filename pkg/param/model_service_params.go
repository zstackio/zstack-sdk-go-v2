// Copyright (c) ZStack.io, Inc.

package param

// UpdateModelServiceDetailParam UpdateModelService详细参数
type UpdateModelServiceDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"name,omitempty"`
	rest string `json:"description,omitempty"`
	rest string `json:"yaml,omitempty"`
	rest int `json:"requestCpu,omitempty"`
	rest int64 `json:"requestMemory,omitempty"`
	rest string `json:"gpuComputeCapability,omitempty"`
	rest string `json:"startCommand,omitempty"`
	rest string `json:"containerCommand,omitempty"`
	rest string `json:"containerArgs,omitempty"`
	rest string `json:"pythonVersion,omitempty"`
	rest string `json:"type,omitempty"`
	rest string `json:"source,omitempty"`
	rest string `json:"framework,omitempty"`
	rest []interface{} `json:"architectureImages,omitempty"`
	rest bool `json:"supportDistributed,omitempty"`
	rest map[string]string `json:"environmentParameters,omitempty"`
	rest map[string]string `json:"startupParameters,omitempty"`
	rest map[string]string `json:"inferenceParams,omitempty"`
	rest string `json:"serviceName,omitempty"`
	rest []int `json:"servicePorts,omitempty"`
	rest string `json:"serviceLivez,omitempty"`
	rest string `json:"serviceReadyz,omitempty"`
	rest int `json:"serviceBootupTime,omitempty"`
	rest string `json:"serviceInstallPath,omitempty"`
	rest string `json:"serviceStartCommand,omitempty"`
	rest map[string]interface{} `json:"vendorToSpecUuidsMap,omitempty"`
}

// UpdateModelServiceParam UpdateModelService请求参数
type UpdateModelServiceParam struct {
	BaseParam
	Params UpdateModelServiceDetailParam `json:"params"` // 详细参数
}

