// Copyright (c) ZStack.io, Inc.

package param

// ChangeTicketStatusDetailParam ChangeTicketStatus详细参数
type ChangeTicketStatusDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"statusEvent" validate:"required"` // 必填
	rest string `json:"comment,omitempty"`
}

// ChangeTicketStatusParam ChangeTicketStatus请求参数
type ChangeTicketStatusParam struct {
	BaseParam
	Params ChangeTicketStatusDetailParam `json:"params"` // 详细参数
}

