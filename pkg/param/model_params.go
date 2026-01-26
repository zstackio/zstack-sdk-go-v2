// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now // avoid unused import

// DeleteModelParamDetail DeleteModel detail param
type DeleteModelParamDetail struct {
	DeleteMode *string `json:"deleteMode,omitempty"`
}

// DeleteModelParam DeleteModel request param
type DeleteModelParam struct {
	BaseParam
	Params DeleteModelParamDetail `json:"deleteModel"`
}
// UpdateModelParamDetail UpdateModel detail param
type UpdateModelParamDetail struct {
	Name string `json:"name,omitempty"`
	InstallPath *string `json:"installPath,omitempty"`
	Description *string `json:"description,omitempty"`
	Parameters *string `json:"parameters,omitempty"`
	ModelCenterUuid *string `json:"modelCenterUuid,omitempty"`
	ModelClassifications []string `json:"modelClassifications,omitempty"`
	RecommendedGpuNum *string `json:"recommendedGpuNum,omitempty"`
	GpuConstraintDescription *string `json:"gpuConstraintDescription,omitempty"`
}

// UpdateModelParam UpdateModel request param
type UpdateModelParam struct {
	BaseParam
	Params UpdateModelParamDetail `json:"updateModel"`
}
// AddModelParamDetail AddModel detail param
type AddModelParamDetail struct {
	Name string `json:"name" validate:"required"`
	InstallPath string `json:"installPath" validate:"required"`
	Description *string `json:"description,omitempty"`
	Parameters *string `json:"parameters,omitempty"`
	Token *string `json:"token,omitempty"`
	ModelCenterUuid string `json:"modelCenterUuid" validate:"required"`
	Logo *string `json:"logo,omitempty"`
	Vendor *string `json:"vendor,omitempty"`
	Introduction *string `json:"introduction,omitempty"`
	Size *int64 `json:"size,omitempty"`
	Version *string `json:"version,omitempty"`
	ModelServiceUuids []string `json:"modelServiceUuids,omitempty"`
	RecommendedGpuNum *string `json:"recommendedGpuNum,omitempty"`
	GpuConstraintDescription *string `json:"gpuConstraintDescription,omitempty"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// AddModelParam AddModel request param
type AddModelParam struct {
	BaseParam
	Params AddModelParamDetail `json:"param"`
}
