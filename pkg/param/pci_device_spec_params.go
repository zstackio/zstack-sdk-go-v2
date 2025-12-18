// Copyright (c) ZStack.io, Inc.

package param

// QueryPciDeviceSpecDetailParam QueryPciDeviceSpec详细参数
type QueryPciDeviceSpecDetailParam struct {
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

// QueryPciDeviceSpecParam QueryPciDeviceSpec请求参数
type QueryPciDeviceSpecParam struct {
	BaseParam
	Params QueryPciDeviceSpecDetailParam `json:"params"` // 详细参数
}

