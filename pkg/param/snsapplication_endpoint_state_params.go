// Copyright (c) ZStack.io, Inc.

package param

// ChangeSNSApplicationEndpointStateDetailParam ChangeSNSApplicationEndpointState详细参数
type ChangeSNSApplicationEndpointStateDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"stateEvent" validate:"required"` // 必填
}

// ChangeSNSApplicationEndpointStateParam ChangeSNSApplicationEndpointState请求参数
type ChangeSNSApplicationEndpointStateParam struct {
	BaseParam
	Params ChangeSNSApplicationEndpointStateDetailParam `json:"params"` // 详细参数
}

