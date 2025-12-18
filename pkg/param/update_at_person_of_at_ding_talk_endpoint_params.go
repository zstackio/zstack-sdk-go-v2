// Copyright (c) ZStack.io, Inc.

package param

// UpdateAtPersonOfAtDingTalkEndpointDetailParam UpdateAtPersonOfAtDingTalkEndpoint detail param
type UpdateAtPersonOfAtDingTalkEndpointDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	EndpointUuid string `json:"endpointUuid" validate:"required"`
	PhoneNumber string `json:"phoneNumber,omitempty"`
	Remark string `json:"remark,omitempty"`
}

// UpdateAtPersonOfAtDingTalkEndpointParam UpdateAtPersonOfAtDingTalkEndpoint request param
type UpdateAtPersonOfAtDingTalkEndpointParam struct {
	BaseParam
	Params UpdateAtPersonOfAtDingTalkEndpointDetailParam `json:"params"`
}
