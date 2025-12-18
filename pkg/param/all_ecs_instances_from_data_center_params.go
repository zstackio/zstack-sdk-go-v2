// Copyright (c) ZStack.io, Inc.

package param

// DeleteAllEcsInstancesFromDataCenterDetailParam DeleteAllEcsInstancesFromDataCenter详细参数
type DeleteAllEcsInstancesFromDataCenterDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"deleteMode,omitempty"`
}

// DeleteAllEcsInstancesFromDataCenterParam DeleteAllEcsInstancesFromDataCenter请求参数
type DeleteAllEcsInstancesFromDataCenterParam struct {
	BaseParam
	Params DeleteAllEcsInstancesFromDataCenterDetailParam `json:"params"` // 详细参数
}

