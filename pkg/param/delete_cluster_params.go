// Copyright (c) ZStack.io, Inc.

package param

// DeleteClusterDetailParam DeleteCluster detail param
type DeleteClusterDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// DeleteClusterParam DeleteCluster request param
type DeleteClusterParam struct {
	BaseParam
	Params DeleteClusterDetailParam `json:"params"`
}
