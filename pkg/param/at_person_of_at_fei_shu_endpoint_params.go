// Copyright (c) ZStack.io, Inc.

package param

// UpdateAtPersonOfAtFeiShuEndpointDetailParam UpdateAtPersonOfAtFeiShuEndpoint详细参数
type UpdateAtPersonOfAtFeiShuEndpointDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"endpointUuid" validate:"required"` // 必填
	rest string `json:"userId,omitempty"`
	rest string `json:"remark,omitempty"`
}

// UpdateAtPersonOfAtFeiShuEndpointParam UpdateAtPersonOfAtFeiShuEndpoint请求参数
type UpdateAtPersonOfAtFeiShuEndpointParam struct {
	BaseParam
	Params UpdateAtPersonOfAtFeiShuEndpointDetailParam `json:"params"` // 详细参数
}

