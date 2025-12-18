// Copyright (c) ZStack.io, Inc.

package param

// DeleteModelServiceInstanceGroupsDetailParam DeleteModelServiceInstanceGroups detail param
type DeleteModelServiceInstanceGroupsDetailParam struct {
	Uuids []string `json:"uuids" validate:"required"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// DeleteModelServiceInstanceGroupsParam DeleteModelServiceInstanceGroups request param
type DeleteModelServiceInstanceGroupsParam struct {
	BaseParam
	Params DeleteModelServiceInstanceGroupsDetailParam `json:"params"`
}
