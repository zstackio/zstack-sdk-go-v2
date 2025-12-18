// Copyright (c) ZStack.io, Inc.

package param

// AddEmailAddressToSNSEmailEndpointDetailParam AddEmailAddressToSNSEmailEndpoint详细参数
type AddEmailAddressToSNSEmailEndpointDetailParam struct {
	rest string `json:"emailAddress" validate:"required"` // 必填
	rest string `json:"endpointUuid" validate:"required"` // 必填
	rest string `json:"resourceUuid,omitempty"`
	rest []string `json:"tagUuids,omitempty"`
}

// AddEmailAddressToSNSEmailEndpointParam AddEmailAddressToSNSEmailEndpoint请求参数
type AddEmailAddressToSNSEmailEndpointParam struct {
	BaseParam
	Params AddEmailAddressToSNSEmailEndpointDetailParam `json:"params"` // 详细参数
}

