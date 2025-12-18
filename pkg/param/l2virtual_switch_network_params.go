// Copyright (c) ZStack.io, Inc.

package param

// QueryL2VirtualSwitchNetworkDetailParam QueryL2VirtualSwitchNetwork详细参数
type QueryL2VirtualSwitchNetworkDetailParam struct {
	rest []interface{} `json:"conditions" validate:"required"` // 必填
	rest int `json:"limit,omitempty"`
	rest int `json:"start,omitempty"`
	rest bool `json:"count,omitempty"`
	rest string `json:"groupBy,omitempty"`
	rest bool `json:"replyWithCount,omitempty"`
	rest string `json:"filterName,omitempty"`
	rest string `json:"sortBy,omitempty"`
	rest string `json:"sortDirection,omitempty"`
	rest []string `json:"fields,omitempty"`
}

// QueryL2VirtualSwitchNetworkParam QueryL2VirtualSwitchNetwork请求参数
type QueryL2VirtualSwitchNetworkParam struct {
	BaseParam
	Params QueryL2VirtualSwitchNetworkDetailParam `json:"params"` // 详细参数
}

