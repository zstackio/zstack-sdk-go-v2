// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// DeleteTicketFlowCollection 删除TicketFlowCollection
func (cli *ZSClient) DeleteTicketFlowCollection(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/tickets/flow-collections/{uuid}", uuid, string(deleteMode))
}

