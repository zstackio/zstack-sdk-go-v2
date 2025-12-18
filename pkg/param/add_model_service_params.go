// Copyright (c) ZStack.io, Inc.

package param

// AddModelServiceDetailParam AddModelService详细参数
type AddModelServiceDetailParam struct {
	rest string `json:"name" validate:"required"` // 必填
	rest string `json:"description,omitempty"`
	rest string `json:"yaml" validate:"required"` // 必填
	rest string `json:"modelCenterUuid" validate:"required"` // 必填
	rest int `json:"requestCpu" validate:"required"` // 必填
	rest int64 `json:"requestMemory" validate:"required"` // 必填
	rest string `json:"gpuComputeCapability,omitempty"`
	rest string `json:"installPath" validate:"required"` // 必填
	rest bool `json:"system,omitempty"`
	rest string `json:"startCommand" validate:"required"` // 必填
	rest string `json:"containerCommand,omitempty"`
	rest string `json:"containerArgs,omitempty"`
	rest string `json:"pythonVersion,omitempty"`
	rest string `json:"condaVersion,omitempty"`
	rest string `json:"zoneUuid,omitempty"`
	rest string `json:"type,omitempty"`
	rest string `json:"source,omitempty"`
	rest string `json:"framework,omitempty"`
	rest []string `json:"modelUuids,omitempty"`
	rest []interface{} `json:"architectureImages,omitempty"`
	rest bool `json:"supportDistributed,omitempty"`
	rest map[string]interface{} `json:"vendorToSpecUuidsMap,omitempty"`
	rest string `json:"resourceUuid,omitempty"`
	rest []string `json:"tagUuids,omitempty"`
}

// AddModelServiceParam AddModelService请求参数
type AddModelServiceParam struct {
	BaseParam
	Params AddModelServiceDetailParam `json:"params"` // 详细参数
}

