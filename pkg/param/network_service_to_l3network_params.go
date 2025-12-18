// Copyright (c) ZStack.io, Inc.

package param

// AttachNetworkServiceToL3NetworkDetailParam AttachNetworkServiceToL3Network详细参数
type AttachNetworkServiceToL3NetworkDetailParam struct {
	rest string `json:"l3NetworkUuid" validate:"required"` // 必填
	rest map[string]interface{} `json:"networkServices" validate:"required"` // 必填
}

// AttachNetworkServiceToL3NetworkParam AttachNetworkServiceToL3Network请求参数
type AttachNetworkServiceToL3NetworkParam struct {
	BaseParam
	Params AttachNetworkServiceToL3NetworkDetailParam `json:"params"` // 详细参数
}

