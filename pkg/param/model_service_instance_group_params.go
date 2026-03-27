// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now() // avoid unused import

// UpdateModelServiceInstanceGroupParamDetail UpdateModelServiceInstanceGroup detail param
type UpdateModelServiceInstanceGroupParamDetail struct {
	Name string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	EnvironmentVariables map[string]string `json:"environmentVariables,omitempty"`
	StartupParameters map[string]string `json:"startupParameters,omitempty"`
	ServiceLivez *string `json:"serviceLivez,omitempty"`
	ServiceReadyz *string `json:"serviceReadyz,omitempty"`
	ServiceBootupTime *int `json:"serviceBootupTime,omitempty"`
}

// UpdateModelServiceInstanceGroupParam UpdateModelServiceInstanceGroup request param
type UpdateModelServiceInstanceGroupParam struct {
	BaseParam
	Params UpdateModelServiceInstanceGroupParamDetail `json:"updateModelServiceInstanceGroup"`
}
// DeleteModelServiceInstanceGroupParamDetail DeleteModelServiceInstanceGroup detail param
type DeleteModelServiceInstanceGroupParamDetail struct {
	DeleteMode *string `json:"deleteMode,omitempty"`
}

// DeleteModelServiceInstanceGroupParam DeleteModelServiceInstanceGroup request param
type DeleteModelServiceInstanceGroupParam struct {
	BaseParam
	Params DeleteModelServiceInstanceGroupParamDetail `json:"deleteModelServiceInstanceGroup"`
}
