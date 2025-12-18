// Copyright (c) ZStack.io, Inc.

package param

// DetachBareMetal2GatewayFromClusterDetailParam DetachBareMetal2GatewayFromCluster detail param
type DetachBareMetal2GatewayFromClusterDetailParam struct {
	ClusterUuid string `json:"clusterUuid" validate:"required"`
	GatewayUuid string `json:"gatewayUuid" validate:"required"`
}

// DetachBareMetal2GatewayFromClusterParam DetachBareMetal2GatewayFromCluster request param
type DetachBareMetal2GatewayFromClusterParam struct {
	BaseParam
	Params DetachBareMetal2GatewayFromClusterDetailParam `json:"params"`
}
