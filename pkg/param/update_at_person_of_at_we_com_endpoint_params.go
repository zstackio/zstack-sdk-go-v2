// Copyright (c) ZStack.io, Inc.

package param

// UpdateAtPersonOfAtWeComEndpointDetailParam UpdateAtPersonOfAtWeComEndpoint detail param
type UpdateAtPersonOfAtWeComEndpointDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	EndpointUuid string `json:"endpointUuid" validate:"required"`
	UserId string `json:"userId,omitempty"`
	Remark string `json:"remark,omitempty"`
}

// UpdateAtPersonOfAtWeComEndpointParam UpdateAtPersonOfAtWeComEndpoint request param
type UpdateAtPersonOfAtWeComEndpointParam struct {
	BaseParam
	Params UpdateAtPersonOfAtWeComEndpointDetailParam `json:"params"`
}
