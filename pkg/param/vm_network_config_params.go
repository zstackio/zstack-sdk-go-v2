// Copyright (c) ZStack.io, Inc.

package param

// UpdateVmNetworkConfigDetailParam UpdateVmNetworkConfig详细参数
type UpdateVmNetworkConfigDetailParam struct {
	rest string `json:"vmInstanceUuid" validate:"required"` // 必填
	rest []string `json:"vmNicUuids" validate:"required"` // 必填
}

// UpdateVmNetworkConfigParam UpdateVmNetworkConfig请求参数
type UpdateVmNetworkConfigParam struct {
	BaseParam
	Params UpdateVmNetworkConfigDetailParam `json:"params"` // 详细参数
}

