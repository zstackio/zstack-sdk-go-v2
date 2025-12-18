// Copyright (c) ZStack.io, Inc.

package param

// UpdateTicketRequestDetailParam UpdateTicketRequest详细参数
type UpdateTicketRequestDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest []interface{} `json:"requests" validate:"required"` // 必填
	rest string `json:"description,omitempty"`
}

// UpdateTicketRequestParam UpdateTicketRequest请求参数
type UpdateTicketRequestParam struct {
	BaseParam
	Params UpdateTicketRequestDetailParam `json:"params"` // 详细参数
}

