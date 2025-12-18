// Copyright (c) ZStack.io, Inc.

package param

// DeleteContainerResourceFromEndpointDetailParam DeleteContainerResourceFromEndpoint详细参数
type DeleteContainerResourceFromEndpointDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
}

// DeleteContainerResourceFromEndpointParam DeleteContainerResourceFromEndpoint请求参数
type DeleteContainerResourceFromEndpointParam struct {
	BaseParam
	Params DeleteContainerResourceFromEndpointDetailParam `json:"params"` // 详细参数
}

