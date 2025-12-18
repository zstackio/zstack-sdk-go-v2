// Copyright (c) ZStack.io, Inc.

package param

// DeleteAllEcsInstancesFromDataCenterDetailParam DeleteAllEcsInstancesFromDataCenter detail param
type DeleteAllEcsInstancesFromDataCenterDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// DeleteAllEcsInstancesFromDataCenterParam DeleteAllEcsInstancesFromDataCenter request param
type DeleteAllEcsInstancesFromDataCenterParam struct {
	BaseParam
	Params DeleteAllEcsInstancesFromDataCenterDetailParam `json:"params"`
}
