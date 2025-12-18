// Copyright (c) ZStack.io, Inc.

package param

// SetIpOnHostNetworkBondingDetailParam SetIpOnHostNetworkBonding详细参数
type SetIpOnHostNetworkBondingDetailParam struct {
	rest string `json:"bondingUuid" validate:"required"` // 必填
	rest string `json:"ipAddress,omitempty"`
	rest string `json:"netmask,omitempty"`
}

// SetIpOnHostNetworkBondingParam SetIpOnHostNetworkBonding请求参数
type SetIpOnHostNetworkBondingParam struct {
	BaseParam
	Params SetIpOnHostNetworkBondingDetailParam `json:"params"` // 详细参数
}

