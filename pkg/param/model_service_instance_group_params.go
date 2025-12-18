// Copyright (c) ZStack.io, Inc.

package param

// UpdateModelServiceInstanceGroupDetailParam UpdateModelServiceInstanceGroup详细参数
type UpdateModelServiceInstanceGroupDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"name,omitempty"`
	rest string `json:"description,omitempty"`
	rest map[string]string `json:"environmentVariables,omitempty"`
	rest map[string]string `json:"startupParameters,omitempty"`
	rest string `json:"serviceLivez,omitempty"`
	rest string `json:"serviceReadyz,omitempty"`
	rest int `json:"serviceBootupTime,omitempty"`
}

// UpdateModelServiceInstanceGroupParam UpdateModelServiceInstanceGroup请求参数
type UpdateModelServiceInstanceGroupParam struct {
	BaseParam
	Params UpdateModelServiceInstanceGroupDetailParam `json:"params"` // 详细参数
}

