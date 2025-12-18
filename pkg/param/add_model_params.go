// Copyright (c) ZStack.io, Inc.

package param

// AddModelDetailParam AddModel detail param
type AddModelDetailParam struct {
	Name string `json:"name" validate:"required"`
	InstallPath string `json:"installPath" validate:"required"`
	Description string `json:"description,omitempty"`
	Parameters string `json:"parameters,omitempty"`
	Token string `json:"token,omitempty"`
	ModelCenterUuid string `json:"modelCenterUuid" validate:"required"`
	Logo string `json:"logo,omitempty"`
	Vendor string `json:"vendor,omitempty"`
	Introduction string `json:"introduction,omitempty"`
	Size int64 `json:"size,omitempty"`
	Version string `json:"version,omitempty"`
	ModelServiceUuids []string `json:"modelServiceUuids,omitempty"`
	RecommendedGpuNum string `json:"recommendedGpuNum,omitempty"`
	GpuConstraintDescription string `json:"gpuConstraintDescription,omitempty"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// AddModelParam AddModel request param
type AddModelParam struct {
	BaseParam
	Params AddModelDetailParam `json:"params"`
}
