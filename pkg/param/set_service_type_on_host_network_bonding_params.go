// Copyright (c) ZStack.io, Inc.

package param

// SetServiceTypeOnHostNetworkBondingDetailParam SetServiceTypeOnHostNetworkBonding详细参数
type SetServiceTypeOnHostNetworkBondingDetailParam struct {
	rest []string `json:"bondingUuids" validate:"required"` // 必填
	rest []int `json:"vlanIds,omitempty"`
	rest []string `json:"serviceTypes,omitempty"`
}

// SetServiceTypeOnHostNetworkBondingParam SetServiceTypeOnHostNetworkBonding请求参数
type SetServiceTypeOnHostNetworkBondingParam struct {
	BaseParam
	Params SetServiceTypeOnHostNetworkBondingDetailParam `json:"params"` // 详细参数
}

