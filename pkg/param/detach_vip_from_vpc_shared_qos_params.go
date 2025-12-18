// Copyright (c) ZStack.io, Inc.

package param

// DetachVipFromVpcSharedQosDetailParam DetachVipFromVpcSharedQos detail param
type DetachVipFromVpcSharedQosDetailParam struct {
	SharedQosUuid string `json:"sharedQosUuid" validate:"required"`
	VipLists []string `json:"vipLists,omitempty"`
	VipUuids []string `json:"vipUuids,omitempty"`
}

// DetachVipFromVpcSharedQosParam DetachVipFromVpcSharedQos request param
type DetachVipFromVpcSharedQosParam struct {
	BaseParam
	Params DetachVipFromVpcSharedQosDetailParam `json:"params"`
}
