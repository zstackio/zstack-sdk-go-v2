// Copyright (c) ZStack.io, Inc.

package param

// RestartModelServiceGroupsDetailParam RestartModelServiceGroups detail param
type RestartModelServiceGroupsDetailParam struct {
	Uuids []string `json:"uuids" validate:"required"`
}

// RestartModelServiceGroupsParam RestartModelServiceGroups request param
type RestartModelServiceGroupsParam struct {
	BaseParam
	Params RestartModelServiceGroupsDetailParam `json:"params"`
}
