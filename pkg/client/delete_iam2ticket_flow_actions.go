// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// DeleteIAM2TicketFlow deletes IAM2TicketFlow
func (cli *ZSClient) DeleteIAM2TicketFlow(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/tickets/flow/{uuid}", uuid, string(deleteMode))
}
