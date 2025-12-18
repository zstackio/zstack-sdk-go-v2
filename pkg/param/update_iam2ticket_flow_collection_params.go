// Copyright (c) ZStack.io, Inc.

package param

// UpdateIAM2TicketFlowCollectionDetailParam UpdateIAM2TicketFlowCollection detail param
type UpdateIAM2TicketFlowCollectionDetailParam struct {
	Flows []interface{} `json:"flows,omitempty"`
	Uuid string `json:"uuid" validate:"required"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	IsDefault bool `json:"isDefault,omitempty"`
}

// UpdateIAM2TicketFlowCollectionParam UpdateIAM2TicketFlowCollection request param
type UpdateIAM2TicketFlowCollectionParam struct {
	BaseParam
	Params UpdateIAM2TicketFlowCollectionDetailParam `json:"params"`
}
