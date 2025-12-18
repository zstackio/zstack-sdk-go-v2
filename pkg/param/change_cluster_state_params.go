// Copyright (c) ZStack.io, Inc.

package param

// ChangeClusterStateDetailParam ChangeClusterState detail param
type ChangeClusterStateDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	StateEvent string `json:"stateEvent" validate:"required"`
}

// ChangeClusterStateParam ChangeClusterState request param
type ChangeClusterStateParam struct {
	BaseParam
	Params ChangeClusterStateDetailParam `json:"params"`
}
