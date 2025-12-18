// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryArchiveTicket queries ArchiveTicket list
func (cli *ZSClient) QueryArchiveTicket(params param.QueryParam) ([]view.ArchiveTicketInventoryView, error) {
	var resp []view.ArchiveTicketInventoryView
	return resp, cli.List("v1/tickets/archives", &params, &resp)
}
