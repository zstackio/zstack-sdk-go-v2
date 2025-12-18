// Copyright (c) ZStack.io, Inc.

package param

// AddIAM2TicketFlowDetailParam AddIAM2TicketFlow detail param
type AddIAM2TicketFlowDetailParam struct {
	ApproverUuid string `json:"approverUuid" validate:"required"`
	ApproverTitle string `json:"approverTitle,omitempty"`
	Name string `json:"name" validate:"required"`
	Description string `json:"description,omitempty"`
	CollectionUuid string `json:"collectionUuid" validate:"required"`
	ParentFlowUuid string `json:"parentFlowUuid,omitempty"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// AddIAM2TicketFlowParam AddIAM2TicketFlow request param
type AddIAM2TicketFlowParam struct {
	BaseParam
	Params AddIAM2TicketFlowDetailParam `json:"params"`
}
