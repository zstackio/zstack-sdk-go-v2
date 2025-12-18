// Copyright (c) ZStack.io, Inc.

package param

// ChangeVmNicNetworkDetailParam ChangeVmNicNetwork详细参数
type ChangeVmNicNetworkDetailParam struct {
	rest string `json:"vmNicUuid" validate:"required"` // 必填
	rest string `json:"destL3NetworkUuid" validate:"required"` // 必填
	rest string `json:"staticIp,omitempty"`
}

// ChangeVmNicNetworkParam ChangeVmNicNetwork请求参数
type ChangeVmNicNetworkParam struct {
	BaseParam
	Params ChangeVmNicNetworkDetailParam `json:"params"` // 详细参数
}

