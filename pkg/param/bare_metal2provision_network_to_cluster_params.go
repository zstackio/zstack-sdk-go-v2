// Copyright (c) ZStack.io, Inc.

package param

// AttachBareMetal2ProvisionNetworkToClusterDetailParam AttachBareMetal2ProvisionNetworkToCluster详细参数
type AttachBareMetal2ProvisionNetworkToClusterDetailParam struct {
	rest string `json:"clusterUuid" validate:"required"` // 必填
	rest string `json:"networkUuid" validate:"required"` // 必填
}

// AttachBareMetal2ProvisionNetworkToClusterParam AttachBareMetal2ProvisionNetworkToCluster请求参数
type AttachBareMetal2ProvisionNetworkToClusterParam struct {
	BaseParam
	Params AttachBareMetal2ProvisionNetworkToClusterDetailParam `json:"params"` // 详细参数
}

