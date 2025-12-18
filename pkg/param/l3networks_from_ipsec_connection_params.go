// Copyright (c) ZStack.io, Inc.

package param

// DetachL3NetworksFromIPsecConnectionDetailParam DetachL3NetworksFromIPsecConnection详细参数
type DetachL3NetworksFromIPsecConnectionDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest []string `json:"l3NetworkUuids" validate:"required"` // 必填
}

// DetachL3NetworksFromIPsecConnectionParam DetachL3NetworksFromIPsecConnection请求参数
type DetachL3NetworksFromIPsecConnectionParam struct {
	BaseParam
	Params DetachL3NetworksFromIPsecConnectionDetailParam `json:"params"` // 详细参数
}

