// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// AddTicketTypesToTicketFlowCollection adds TicketTypesToTicketFlowCollection
func (cli *ZSClient) AddTicketTypesToTicketFlowCollection(params param.AddTicketTypesToTicketFlowCollectionParam) (*view.AddTicketTypesToTicketFlowCollectionEventView, error) {
	resp := view.AddTicketTypesToTicketFlowCollectionEventView{}
	if err := cli.Post("v1/tickets/flow-collections/{ticketFlowCollectionUuid}/ticket-types", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
