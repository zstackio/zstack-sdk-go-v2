// Copyright (c) ZStack.io, Inc.

package param

// ChangeBareMetal2GatewayClusterDetailParam ChangeBareMetal2GatewayCluster detail param
type ChangeBareMetal2GatewayClusterDetailParam struct {
	GatewayUuid string `json:"gatewayUuid" validate:"required"`
	ClusterUuid string `json:"clusterUuid" validate:"required"`
}

// ChangeBareMetal2GatewayClusterParam ChangeBareMetal2GatewayCluster request param
type ChangeBareMetal2GatewayClusterParam struct {
	BaseParam
	Params ChangeBareMetal2GatewayClusterDetailParam `json:"params"`
}
