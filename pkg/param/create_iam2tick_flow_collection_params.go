// Copyright (c) ZStack.io, Inc.

package param

// CreateIAM2TickFlowCollectionDetailParam CreateIAM2TickFlowCollection detail param
type CreateIAM2TickFlowCollectionDetailParam struct {
	Flows []IAM2FlowStructParam `json:"flows,omitempty"`
	ProjectUuid string `json:"projectUuid" validate:"required"`
	Name string `json:"name" validate:"required"`
	Description string `json:"description,omitempty"`
	IsDefault bool `json:"isDefault,omitempty"`
	TicketTypeUuids []string `json:"ticketTypeUuids,omitempty"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateIAM2TickFlowCollectionParam CreateIAM2TickFlowCollection request param
type CreateIAM2TickFlowCollectionParam struct {
	BaseParam
	Params CreateIAM2TickFlowCollectionDetailParam `json:"params"`
}
