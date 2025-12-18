// Copyright (c) ZStack.io, Inc.

package param

// CreateHostNetworkServiceTypeDetailParam CreateHostNetworkServiceType详细参数
type CreateHostNetworkServiceTypeDetailParam struct {
	rest string `json:"serviceType" validate:"required"` // 必填
	rest bool `json:"system,omitempty"`
	rest string `json:"resourceUuid,omitempty"`
	rest []string `json:"tagUuids,omitempty"`
}

// CreateHostNetworkServiceTypeParam CreateHostNetworkServiceType请求参数
type CreateHostNetworkServiceTypeParam struct {
	BaseParam
	Params CreateHostNetworkServiceTypeDetailParam `json:"params"` // 详细参数
}

