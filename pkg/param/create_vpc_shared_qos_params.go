// Copyright (c) ZStack.io, Inc.

package param

// CreateVpcSharedQosDetailParam CreateVpcSharedQos detail param
type CreateVpcSharedQosDetailParam struct {
	Name string `json:"name" validate:"required"`
	Description string `json:"description,omitempty"`
	VpcUuid string `json:"vpcUuid" validate:"required"`
	L3NetworkUuid string `json:"l3NetworkUuid" validate:"required"`
	Bandwidth int64 `json:"bandwidth,omitempty"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateVpcSharedQosParam CreateVpcSharedQos request param
type CreateVpcSharedQosParam struct {
	BaseParam
	Params CreateVpcSharedQosDetailParam `json:"params"`
}
