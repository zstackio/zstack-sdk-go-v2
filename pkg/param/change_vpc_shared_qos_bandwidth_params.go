// Copyright (c) ZStack.io, Inc.

package param

// ChangeVpcSharedQosBandwidthDetailParam ChangeVpcSharedQosBandwidth detail param
type ChangeVpcSharedQosBandwidthDetailParam struct {
	SharedQosUuid string `json:"sharedQosUuid" validate:"required"`
	Bandwidth int64 `json:"bandwidth" validate:"required"`
}

// ChangeVpcSharedQosBandwidthParam ChangeVpcSharedQosBandwidth request param
type ChangeVpcSharedQosBandwidthParam struct {
	BaseParam
	Params ChangeVpcSharedQosBandwidthDetailParam `json:"params"`
}
