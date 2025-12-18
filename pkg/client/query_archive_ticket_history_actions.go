// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryArchiveTicketHistory queries ArchiveTicketHistory list
func (cli *ZSClient) QueryArchiveTicketHistory(params param.QueryParam) ([]view.ArchiveTicketStatusHistoryInventoryView, error) {
	var resp []view.ArchiveTicketStatusHistoryInventoryView
	return resp, cli.List("v1/tickets/histories/archives", &params, &resp)
}
