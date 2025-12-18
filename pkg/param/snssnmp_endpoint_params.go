// Copyright (c) ZStack.io, Inc.

package param

// CreateSNSSnmpEndpointDetailParam CreateSNSSnmpEndpoint详细参数
type CreateSNSSnmpEndpointDetailParam struct {
	rest string `json:"name" validate:"required"` // 必填
	rest string `json:"description,omitempty"`
	rest string `json:"platformUuid,omitempty"`
	rest string `json:"resourceUuid,omitempty"`
	rest []string `json:"tagUuids,omitempty"`
}

// CreateSNSSnmpEndpointParam CreateSNSSnmpEndpoint请求参数
type CreateSNSSnmpEndpointParam struct {
	BaseParam
	Params CreateSNSSnmpEndpointDetailParam `json:"params"` // 详细参数
}

