// Copyright (c) ZStack.io, Inc.

package param

// GetAvailableVpcL3NetworkDetailParam GetAvailableVpcL3Network详细参数
type GetAvailableVpcL3NetworkDetailParam struct {
	rest string `json:"zoneUuid" validate:"required"` // 必填
}

// GetAvailableVpcL3NetworkParam GetAvailableVpcL3Network请求参数
type GetAvailableVpcL3NetworkParam struct {
	BaseParam
	Params GetAvailableVpcL3NetworkDetailParam `json:"params"` // 详细参数
}

