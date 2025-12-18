// Copyright (c) ZStack.io, Inc.

package param

// DeleteModelServiceInstanceGroupsDetailParam DeleteModelServiceInstanceGroups详细参数
type DeleteModelServiceInstanceGroupsDetailParam struct {
	rest []string `json:"uuids" validate:"required"` // 必填
	rest string `json:"deleteMode,omitempty"`
}

// DeleteModelServiceInstanceGroupsParam DeleteModelServiceInstanceGroups请求参数
type DeleteModelServiceInstanceGroupsParam struct {
	BaseParam
	Params DeleteModelServiceInstanceGroupsDetailParam `json:"params"` // 详细参数
}

