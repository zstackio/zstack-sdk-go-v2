// Copyright (c) ZStack.io, Inc.

package param

// AttachBareMetal2ProvisionNetworkToClusterDetailParam AttachBareMetal2ProvisionNetworkToCluster detail param
type AttachBareMetal2ProvisionNetworkToClusterDetailParam struct {
	ClusterUuid string `json:"clusterUuid" validate:"required"`
	NetworkUuid string `json:"networkUuid" validate:"required"`
}

// AttachBareMetal2ProvisionNetworkToClusterParam AttachBareMetal2ProvisionNetworkToCluster request param
type AttachBareMetal2ProvisionNetworkToClusterParam struct {
	BaseParam
	Params AttachBareMetal2ProvisionNetworkToClusterDetailParam `json:"params"`
}
