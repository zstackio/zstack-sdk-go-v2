// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now // avoid unused import

// DeleteTicketFlowCollectionParamDetail DeleteTicketFlowCollection detail param
type DeleteTicketFlowCollectionParamDetail struct {
	Uuid string `json:"uuid" validate:"required"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// DeleteTicketFlowCollectionParam DeleteTicketFlowCollection request param
type DeleteTicketFlowCollectionParam struct {
	BaseParam
	Params DeleteTicketFlowCollectionParamDetail `json:"deleteTicketFlowCollection"`
}
