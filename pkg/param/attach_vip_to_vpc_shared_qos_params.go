// Copyright (c) ZStack.io, Inc.

package param

// AttachVipToVpcSharedQosDetailParam AttachVipToVpcSharedQos detail param
type AttachVipToVpcSharedQosDetailParam struct {
	SharedQosUuid string `json:"sharedQosUuid" validate:"required"`
	VipLists []string `json:"vipLists,omitempty"`
	VipUuids []string `json:"vipUuids,omitempty"`
}

// AttachVipToVpcSharedQosParam AttachVipToVpcSharedQos request param
type AttachVipToVpcSharedQosParam struct {
	BaseParam
	Params AttachVipToVpcSharedQosDetailParam `json:"params"`
}
