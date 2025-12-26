// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QueryArchiveTicketHistory queries ArchiveTicketHistory list
func (cli *ZSClient) QueryArchiveTicketHistory(params *param.QueryParam) ([]view.ArchiveTicketStatusHistoryInventoryView, error) {
	var resp []view.ArchiveTicketStatusHistoryInventoryView
	return resp, cli.List("v1/tickets/histories/archives", params, &resp)
}
