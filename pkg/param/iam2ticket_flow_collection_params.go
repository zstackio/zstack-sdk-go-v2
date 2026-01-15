// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now // avoid unused import

// UpdateIAM2TicketFlowCollectionParamDetail UpdateIAM2TicketFlowCollection detail param
type UpdateIAM2TicketFlowCollectionParamDetail struct {
	Flows []UpdateIAM2TicketFlowCollection_IAM2FlowStructParam `json:"flows,omitempty"`
	Uuid string `json:"uuid" validate:"required"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	IsDefault bool `json:"isDefault,omitempty"`
}

// UpdateIAM2TicketFlowCollectionParam UpdateIAM2TicketFlowCollection request param
type UpdateIAM2TicketFlowCollectionParam struct {
	BaseParam
	Params UpdateIAM2TicketFlowCollectionParamDetail `json:"updateIAM2TicketFlowCollection"`
}
