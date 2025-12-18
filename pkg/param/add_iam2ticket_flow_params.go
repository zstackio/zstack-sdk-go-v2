// Copyright (c) ZStack.io, Inc.

package param

// AddIAM2TicketFlowDetailParam AddIAM2TicketFlow详细参数
type AddIAM2TicketFlowDetailParam struct {
	rest string `json:"approverUuid" validate:"required"` // 必填
	rest string `json:"approverTitle,omitempty"`
	rest string `json:"name" validate:"required"` // 必填
	rest string `json:"description,omitempty"`
	rest string `json:"collectionUuid" validate:"required"` // 必填
	rest string `json:"parentFlowUuid,omitempty"`
	rest string `json:"resourceUuid,omitempty"`
	rest []string `json:"tagUuids,omitempty"`
}

// AddIAM2TicketFlowParam AddIAM2TicketFlow请求参数
type AddIAM2TicketFlowParam struct {
	BaseParam
	Params AddIAM2TicketFlowDetailParam `json:"params"` // 详细参数
}

