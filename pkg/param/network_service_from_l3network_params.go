// Copyright (c) ZStack.io, Inc.

package param

// DetachNetworkServiceFromL3NetworkDetailParam DetachNetworkServiceFromL3Network详细参数
type DetachNetworkServiceFromL3NetworkDetailParam struct {
	rest string `json:"l3NetworkUuid" validate:"required"` // 必填
	rest map[string]interface{} `json:"networkServices,omitempty"`
	rest string `json:"service,omitempty"`
}

// DetachNetworkServiceFromL3NetworkParam DetachNetworkServiceFromL3Network请求参数
type DetachNetworkServiceFromL3NetworkParam struct {
	BaseParam
	Params DetachNetworkServiceFromL3NetworkDetailParam `json:"params"` // 详细参数
}

