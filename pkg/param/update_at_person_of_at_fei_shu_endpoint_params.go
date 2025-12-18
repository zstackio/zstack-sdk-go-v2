// Copyright (c) ZStack.io, Inc.

package param

// UpdateAtPersonOfAtFeiShuEndpointDetailParam UpdateAtPersonOfAtFeiShuEndpoint detail param
type UpdateAtPersonOfAtFeiShuEndpointDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	EndpointUuid string `json:"endpointUuid" validate:"required"`
	UserId string `json:"userId,omitempty"`
	Remark string `json:"remark,omitempty"`
}

// UpdateAtPersonOfAtFeiShuEndpointParam UpdateAtPersonOfAtFeiShuEndpoint request param
type UpdateAtPersonOfAtFeiShuEndpointParam struct {
	BaseParam
	Params UpdateAtPersonOfAtFeiShuEndpointDetailParam `json:"params"`
}
