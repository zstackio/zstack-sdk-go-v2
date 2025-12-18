// Copyright (c) ZStack.io, Inc.

package param

// UpdateClusterDetailParam UpdateCluster detail param
type UpdateClusterDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

// UpdateClusterParam UpdateCluster request param
type UpdateClusterParam struct {
	BaseParam
	Params UpdateClusterDetailParam `json:"params"`
}
