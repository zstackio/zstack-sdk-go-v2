// Copyright (c) ZStack.io, Inc.

package param

// DetachBareMetal2ProvisionNetworkFromClusterDetailParam DetachBareMetal2ProvisionNetworkFromCluster详细参数
type DetachBareMetal2ProvisionNetworkFromClusterDetailParam struct {
	rest string `json:"clusterUuid" validate:"required"` // 必填
	rest string `json:"networkUuid" validate:"required"` // 必填
}

// DetachBareMetal2ProvisionNetworkFromClusterParam DetachBareMetal2ProvisionNetworkFromCluster请求参数
type DetachBareMetal2ProvisionNetworkFromClusterParam struct {
	BaseParam
	Params DetachBareMetal2ProvisionNetworkFromClusterDetailParam `json:"params"` // 详细参数
}

