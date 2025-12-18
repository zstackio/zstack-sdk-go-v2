// Copyright (c) ZStack.io, Inc.

package param

// UpdateAtPersonOfAtDingTalkEndpointDetailParam UpdateAtPersonOfAtDingTalkEndpoint详细参数
type UpdateAtPersonOfAtDingTalkEndpointDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"endpointUuid" validate:"required"` // 必填
	rest string `json:"phoneNumber,omitempty"`
	rest string `json:"remark,omitempty"`
}

// UpdateAtPersonOfAtDingTalkEndpointParam UpdateAtPersonOfAtDingTalkEndpoint请求参数
type UpdateAtPersonOfAtDingTalkEndpointParam struct {
	BaseParam
	Params UpdateAtPersonOfAtDingTalkEndpointDetailParam `json:"params"` // 详细参数
}

