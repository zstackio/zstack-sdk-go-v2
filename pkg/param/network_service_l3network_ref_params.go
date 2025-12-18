// Copyright (c) ZStack.io, Inc.

package param

// QueryNetworkServiceL3NetworkRefDetailParam QueryNetworkServiceL3NetworkRef详细参数
type QueryNetworkServiceL3NetworkRefDetailParam struct {
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

// QueryNetworkServiceL3NetworkRefParam QueryNetworkServiceL3NetworkRef请求参数
type QueryNetworkServiceL3NetworkRefParam struct {
	BaseParam
	Params QueryNetworkServiceL3NetworkRefDetailParam `json:"params"` // 详细参数
}

