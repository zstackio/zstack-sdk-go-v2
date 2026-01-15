// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now // avoid unused import

// UpdateIAM2TicketFlowParamDetail UpdateIAM2TicketFlow detail param
type UpdateIAM2TicketFlowParamDetail struct {
	Uuid string `json:"uuid" validate:"required"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	ApproverUuid string `json:"approverUuid,omitempty"`
	ApproverTitle string `json:"approverTitle,omitempty"`
}

// UpdateIAM2TicketFlowParam UpdateIAM2TicketFlow request param
type UpdateIAM2TicketFlowParam struct {
	BaseParam
	Params UpdateIAM2TicketFlowParamDetail `json:"updateIAM2TicketFlow"`
}
// DeleteIAM2TicketFlowParamDetail DeleteIAM2TicketFlow detail param
type DeleteIAM2TicketFlowParamDetail struct {
	Uuid string `json:"uuid" validate:"required"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// DeleteIAM2TicketFlowParam DeleteIAM2TicketFlow request param
type DeleteIAM2TicketFlowParam struct {
	BaseParam
	Params DeleteIAM2TicketFlowParamDetail `json:"deleteIAM2TicketFlow"`
}
// AddIAM2TicketFlowParamDetail AddIAM2TicketFlow detail param
type AddIAM2TicketFlowParamDetail struct {
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
	Params AddIAM2TicketFlowParamDetail `json:"addIAM2TicketFlow"`
}
