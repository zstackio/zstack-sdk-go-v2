// Copyright (c) ZStack.io, Inc.

package param

// UpdateModelDetailParam UpdateModel detail param
type UpdateModelDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	Name string `json:"name,omitempty"`
	InstallPath string `json:"installPath,omitempty"`
	Description string `json:"description,omitempty"`
	Parameters string `json:"parameters,omitempty"`
	ModelCenterUuid string `json:"modelCenterUuid,omitempty"`
	ModelClassifications []string `json:"modelClassifications,omitempty"`
	RecommendedGpuNum string `json:"recommendedGpuNum,omitempty"`
	GpuConstraintDescription string `json:"gpuConstraintDescription,omitempty"`
}

// UpdateModelParam UpdateModel request param
type UpdateModelParam struct {
	BaseParam
	Params UpdateModelDetailParam `json:"params"`
}
