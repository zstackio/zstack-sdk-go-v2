// Copyright (c) ZStack.io, Inc.

package param

// CreateSNSSnmpEndpointDetailParam CreateSNSSnmpEndpoint detail param
type CreateSNSSnmpEndpointDetailParam struct {
	Name string `json:"name" validate:"required"`
	Description string `json:"description,omitempty"`
	PlatformUuid string `json:"platformUuid,omitempty"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateSNSSnmpEndpointParam CreateSNSSnmpEndpoint request param
type CreateSNSSnmpEndpointParam struct {
	BaseParam
	Params CreateSNSSnmpEndpointDetailParam `json:"params"`
}
