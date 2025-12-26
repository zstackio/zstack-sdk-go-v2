// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// RemoveTicketTypesFromTicketFlowCollection removes TicketTypesFromTicketFlowCollection
func (cli *ZSClient) RemoveTicketTypesFromTicketFlowCollection(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/tickets/flow-collections/{ticketFlowCollectionUuid}/ticket-types", uuid, string(deleteMode))
}
