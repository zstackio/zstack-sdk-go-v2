// Copyright (c) ZStack.io, Inc.

package param

// DetachBareMetal2ProvisionNetworkFromClusterDetailParam DetachBareMetal2ProvisionNetworkFromCluster detail param
type DetachBareMetal2ProvisionNetworkFromClusterDetailParam struct {
	ClusterUuid string `json:"clusterUuid" validate:"required"`
	NetworkUuid string `json:"networkUuid" validate:"required"`
}

// DetachBareMetal2ProvisionNetworkFromClusterParam DetachBareMetal2ProvisionNetworkFromCluster request param
type DetachBareMetal2ProvisionNetworkFromClusterParam struct {
	BaseParam
	Params DetachBareMetal2ProvisionNetworkFromClusterDetailParam `json:"params"`
}
