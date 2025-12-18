// Copyright (c) ZStack.io, Inc.

package param

// AddEmailAddressToSNSEmailEndpointDetailParam AddEmailAddressToSNSEmailEndpoint detail param
type AddEmailAddressToSNSEmailEndpointDetailParam struct {
	EmailAddress string `json:"emailAddress" validate:"required"`
	EndpointUuid string `json:"endpointUuid" validate:"required"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// AddEmailAddressToSNSEmailEndpointParam AddEmailAddressToSNSEmailEndpoint request param
type AddEmailAddressToSNSEmailEndpointParam struct {
	BaseParam
	Params AddEmailAddressToSNSEmailEndpointDetailParam `json:"params"`
}
