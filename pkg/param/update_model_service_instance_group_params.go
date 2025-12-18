// Copyright (c) ZStack.io, Inc.

package param

// UpdateModelServiceInstanceGroupDetailParam UpdateModelServiceInstanceGroup detail param
type UpdateModelServiceInstanceGroupDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	EnvironmentVariables map[string]string `json:"environmentVariables,omitempty"`
	StartupParameters map[string]string `json:"startupParameters,omitempty"`
	ServiceLivez string `json:"serviceLivez,omitempty"`
	ServiceReadyz string `json:"serviceReadyz,omitempty"`
	ServiceBootupTime int `json:"serviceBootupTime,omitempty"`
}

// UpdateModelServiceInstanceGroupParam UpdateModelServiceInstanceGroup request param
type UpdateModelServiceInstanceGroupParam struct {
	BaseParam
	Params UpdateModelServiceInstanceGroupDetailParam `json:"params"`
}
