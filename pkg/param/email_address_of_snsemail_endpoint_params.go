// Copyright (c) ZStack.io, Inc.

package param

// DeleteEmailAddressOfSNSEmailEndpointDetailParam DeleteEmailAddressOfSNSEmailEndpoint详细参数
type DeleteEmailAddressOfSNSEmailEndpointDetailParam struct {
	rest string `json:"emailAddressUuid" validate:"required"` // 必填
	rest string `json:"endpointUuid" validate:"required"` // 必填
}

// DeleteEmailAddressOfSNSEmailEndpointParam DeleteEmailAddressOfSNSEmailEndpoint请求参数
type DeleteEmailAddressOfSNSEmailEndpointParam struct {
	BaseParam
	Params DeleteEmailAddressOfSNSEmailEndpointDetailParam `json:"params"` // 详细参数
}

