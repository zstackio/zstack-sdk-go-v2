// Copyright (c) ZStack.io, Inc.

package param

// RestartModelServiceGroupsDetailParam RestartModelServiceGroups详细参数
type RestartModelServiceGroupsDetailParam struct {
	rest []string `json:"uuids" validate:"required"` // 必填
}

// RestartModelServiceGroupsParam RestartModelServiceGroups请求参数
type RestartModelServiceGroupsParam struct {
	BaseParam
	Params RestartModelServiceGroupsDetailParam `json:"params"` // 详细参数
}

