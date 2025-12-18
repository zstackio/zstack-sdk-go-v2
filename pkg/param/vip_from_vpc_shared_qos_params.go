// Copyright (c) ZStack.io, Inc.

package param

// DetachVipFromVpcSharedQosDetailParam DetachVipFromVpcSharedQos详细参数
type DetachVipFromVpcSharedQosDetailParam struct {
	rest string `json:"sharedQosUuid" validate:"required"` // 必填
	rest []string `json:"vipLists,omitempty"`
	rest []string `json:"vipUuids,omitempty"`
}

// DetachVipFromVpcSharedQosParam DetachVipFromVpcSharedQos请求参数
type DetachVipFromVpcSharedQosParam struct {
	BaseParam
	Params DetachVipFromVpcSharedQosDetailParam `json:"params"` // 详细参数
}

