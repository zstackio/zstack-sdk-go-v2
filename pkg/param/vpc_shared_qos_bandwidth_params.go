// Copyright (c) ZStack.io, Inc.

package param

// ChangeVpcSharedQosBandwidthDetailParam ChangeVpcSharedQosBandwidth详细参数
type ChangeVpcSharedQosBandwidthDetailParam struct {
	rest string `json:"sharedQosUuid" validate:"required"` // 必填
	rest int64 `json:"bandwidth" validate:"required"` // 必填
}

// ChangeVpcSharedQosBandwidthParam ChangeVpcSharedQosBandwidth请求参数
type ChangeVpcSharedQosBandwidthParam struct {
	BaseParam
	Params ChangeVpcSharedQosBandwidthDetailParam `json:"params"` // 详细参数
}

