// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// AddTicketTypesToTicketFlowCollection adds TicketTypesToTicketFlowCollection
func (cli *ZSClient) AddTicketTypesToTicketFlowCollection(params param.AddTicketTypesToTicketFlowCollectionParam) (*view.AddTicketTypesToTicketFlowCollectionEventView, error) {
	resp := view.AddTicketTypesToTicketFlowCollectionEventView{}
	if err := cli.Post("v1/tickets/flow-collections/{ticketFlowCollectionUuid}/ticket-types", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
