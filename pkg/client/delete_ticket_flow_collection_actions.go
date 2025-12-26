// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// DeleteTicketFlowCollection deletes TicketFlowCollection
func (cli *ZSClient) DeleteTicketFlowCollection(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/tickets/flow-collections/{uuid}", uuid, string(deleteMode))
}
