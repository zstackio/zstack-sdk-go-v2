// Copyright (c) ZStack.io, Inc.

package param

// RemoveDnsFromL3NetworkDetailParam RemoveDnsFromL3Network详细参数
type RemoveDnsFromL3NetworkDetailParam struct {
	rest string `json:"l3NetworkUuid" validate:"required"` // 必填
	rest string `json:"dns" validate:"required"` // 必填
}

// RemoveDnsFromL3NetworkParam RemoveDnsFromL3Network请求参数
type RemoveDnsFromL3NetworkParam struct {
	BaseParam
	Params RemoveDnsFromL3NetworkDetailParam `json:"params"` // 详细参数
}

