// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// RemoveTicketTypesFromTicketFlowCollection 操作RemoveTicketTypesFromTicketFlowCollection
func (cli *ZSClient) RemoveTicketTypesFromTicketFlowCollection(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/tickets/flow-collections/{ticketFlowCollectionUuid}/ticket-types", uuid, string(deleteMode))
}

