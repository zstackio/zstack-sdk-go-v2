// Copyright (c) ZStack.io, Inc.

package param

// DeleteModelServiceInstanceGroupDetailParam DeleteModelServiceInstanceGroup detail param
type DeleteModelServiceInstanceGroupDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// DeleteModelServiceInstanceGroupParam DeleteModelServiceInstanceGroup request param
type DeleteModelServiceInstanceGroupParam struct {
	BaseParam
	Params DeleteModelServiceInstanceGroupDetailParam `json:"params"`
}
