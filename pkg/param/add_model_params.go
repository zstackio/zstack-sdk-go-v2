// Copyright (c) ZStack.io, Inc.

package param

// AddModelDetailParam AddModel详细参数
type AddModelDetailParam struct {
	rest string `json:"name" validate:"required"` // 必填
	rest string `json:"installPath" validate:"required"` // 必填
	rest string `json:"description,omitempty"`
	rest string `json:"parameters,omitempty"`
	rest string `json:"token,omitempty"`
	rest string `json:"modelCenterUuid" validate:"required"` // 必填
	rest string `json:"logo,omitempty"`
	rest string `json:"vendor,omitempty"`
	rest string `json:"introduction,omitempty"`
	rest int64 `json:"size,omitempty"`
	rest string `json:"version,omitempty"`
	rest []string `json:"modelServiceUuids,omitempty"`
	rest string `json:"recommendedGpuNum,omitempty"`
	rest string `json:"gpuConstraintDescription,omitempty"`
	rest string `json:"resourceUuid,omitempty"`
	rest []string `json:"tagUuids,omitempty"`
}

// AddModelParam AddModel请求参数
type AddModelParam struct {
	BaseParam
	Params AddModelDetailParam `json:"params"` // 详细参数
}

