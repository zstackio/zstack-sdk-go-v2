// Copyright (c) ZStack.io, Inc.

package param

// CreateTicketDetailParam CreateTicket详细参数
type CreateTicketDetailParam struct {
	rest string `json:"name" validate:"required"` // 必填
	rest string `json:"description,omitempty"`
	rest []interface{} `json:"requests" validate:"required"` // 必填
	rest string `json:"flowCollectionUuid,omitempty"`
	rest string `json:"accountSystemType" validate:"required"` // 必填
	rest interface{} `json:"accountSystemContext" validate:"required"` // 必填
	rest string `json:"resourceUuid,omitempty"`
	rest []string `json:"tagUuids,omitempty"`
}

// CreateTicketParam CreateTicket请求参数
type CreateTicketParam struct {
	BaseParam
	Params CreateTicketDetailParam `json:"params"` // 详细参数
}

