// Copyright (c) ZStack.io, Inc.

package param

// QueryBareMetal2ChassisOfferingDetailParam QueryBareMetal2ChassisOffering详细参数
type QueryBareMetal2ChassisOfferingDetailParam struct {
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

// QueryBareMetal2ChassisOfferingParam QueryBareMetal2ChassisOffering请求参数
type QueryBareMetal2ChassisOfferingParam struct {
	BaseParam
	Params QueryBareMetal2ChassisOfferingDetailParam `json:"params"` // 详细参数
}

