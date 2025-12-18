// Copyright (c) ZStack.io, Inc.

package param

// AttachBareMetal2GatewayToClusterDetailParam AttachBareMetal2GatewayToCluster detail param
type AttachBareMetal2GatewayToClusterDetailParam struct {
	ClusterUuid string `json:"clusterUuid" validate:"required"`
	GatewayUuid string `json:"gatewayUuid" validate:"required"`
}

// AttachBareMetal2GatewayToClusterParam AttachBareMetal2GatewayToCluster request param
type AttachBareMetal2GatewayToClusterParam struct {
	BaseParam
	Params AttachBareMetal2GatewayToClusterDetailParam `json:"params"`
}
