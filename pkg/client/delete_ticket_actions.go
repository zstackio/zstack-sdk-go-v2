// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// DeleteTicket deletes Ticket
func (cli *ZSClient) DeleteTicket(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/tickets/{uuid}", uuid, string(deleteMode))
}
