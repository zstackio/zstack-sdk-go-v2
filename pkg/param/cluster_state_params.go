// Copyright (c) ZStack.io, Inc.

package param

// ChangeClusterStateDetailParam ChangeClusterState详细参数
type ChangeClusterStateDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"stateEvent" validate:"required"` // 必填
}

// ChangeClusterStateParam ChangeClusterState请求参数
type ChangeClusterStateParam struct {
	BaseParam
	Params ChangeClusterStateDetailParam `json:"params"` // 详细参数
}

