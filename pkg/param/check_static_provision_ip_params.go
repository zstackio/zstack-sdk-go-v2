// Copyright (c) ZStack.io, Inc.

package param

// CheckStaticProvisionIpDetailParam CheckStaticProvisionIp detail param
type CheckStaticProvisionIpDetailParam struct {
	ClusterUuid string `json:"clusterUuid" validate:"required"`
	ProvisionIp string `json:"provisionIp" validate:"required"`
}

// CheckStaticProvisionIpParam CheckStaticProvisionIp request param
type CheckStaticProvisionIpParam struct {
	BaseParam
	Params CheckStaticProvisionIpDetailParam `json:"params"`
}
