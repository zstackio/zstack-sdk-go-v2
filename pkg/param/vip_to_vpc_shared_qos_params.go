// Copyright (c) ZStack.io, Inc.

package param

// AttachVipToVpcSharedQosDetailParam AttachVipToVpcSharedQos详细参数
type AttachVipToVpcSharedQosDetailParam struct {
	rest string `json:"sharedQosUuid" validate:"required"` // 必填
	rest []string `json:"vipLists,omitempty"`
	rest []string `json:"vipUuids,omitempty"`
}

// AttachVipToVpcSharedQosParam AttachVipToVpcSharedQos请求参数
type AttachVipToVpcSharedQosParam struct {
	BaseParam
	Params AttachVipToVpcSharedQosDetailParam `json:"params"` // 详细参数
}

