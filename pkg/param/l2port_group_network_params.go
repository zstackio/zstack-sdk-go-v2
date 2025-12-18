// Copyright (c) ZStack.io, Inc.

package param

// QueryL2PortGroupNetworkDetailParam QueryL2PortGroupNetwork详细参数
type QueryL2PortGroupNetworkDetailParam struct {
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

// QueryL2PortGroupNetworkParam QueryL2PortGroupNetwork请求参数
type QueryL2PortGroupNetworkParam struct {
	BaseParam
	Params QueryL2PortGroupNetworkDetailParam `json:"params"` // 详细参数
}

