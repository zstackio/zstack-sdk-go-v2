// Copyright (c) ZStack.io, Inc.

package param

// UpdateEmailAddressOfSNSEmailEndpointDetailParam UpdateEmailAddressOfSNSEmailEndpoint detail param
type UpdateEmailAddressOfSNSEmailEndpointDetailParam struct {
	EmailAddressUuid string `json:"emailAddressUuid" validate:"required"`
	EndpointUuid string `json:"endpointUuid" validate:"required"`
	EmailAddress string `json:"emailAddress" validate:"required"`
}

// UpdateEmailAddressOfSNSEmailEndpointParam UpdateEmailAddressOfSNSEmailEndpoint request param
type UpdateEmailAddressOfSNSEmailEndpointParam struct {
	BaseParam
	Params UpdateEmailAddressOfSNSEmailEndpointDetailParam `json:"params"`
}
