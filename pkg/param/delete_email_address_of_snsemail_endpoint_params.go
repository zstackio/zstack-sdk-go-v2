// Copyright (c) ZStack.io, Inc.

package param

// DeleteEmailAddressOfSNSEmailEndpointDetailParam DeleteEmailAddressOfSNSEmailEndpoint detail param
type DeleteEmailAddressOfSNSEmailEndpointDetailParam struct {
	EmailAddressUuid string `json:"emailAddressUuid" validate:"required"`
	EndpointUuid string `json:"endpointUuid" validate:"required"`
}

// DeleteEmailAddressOfSNSEmailEndpointParam DeleteEmailAddressOfSNSEmailEndpoint request param
type DeleteEmailAddressOfSNSEmailEndpointParam struct {
	BaseParam
	Params DeleteEmailAddressOfSNSEmailEndpointDetailParam `json:"params"`
}
