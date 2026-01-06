// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryArchiveTicket queries ArchiveTicket list
func (cli *ZSClient) QueryArchiveTicket(params *param.QueryParam) ([]view.ArchiveTicketInventoryView, error) {
	var resp []view.ArchiveTicketInventoryView
	return resp, cli.List("v1/tickets/archives", params, &resp)
}
