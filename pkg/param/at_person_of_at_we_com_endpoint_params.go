// Copyright (c) ZStack.io, Inc.

package param

// UpdateAtPersonOfAtWeComEndpointDetailParam UpdateAtPersonOfAtWeComEndpoint详细参数
type UpdateAtPersonOfAtWeComEndpointDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"endpointUuid" validate:"required"` // 必填
	rest string `json:"userId,omitempty"`
	rest string `json:"remark,omitempty"`
}

// UpdateAtPersonOfAtWeComEndpointParam UpdateAtPersonOfAtWeComEndpoint请求参数
type UpdateAtPersonOfAtWeComEndpointParam struct {
	BaseParam
	Params UpdateAtPersonOfAtWeComEndpointDetailParam `json:"params"` // 详细参数
}

