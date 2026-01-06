// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryTicketType queries TicketType list
func (cli *ZSClient) QueryTicketType(params *param.QueryParam) ([]view.TicketTypeInventoryView, error) {
	var resp []view.TicketTypeInventoryView
	return resp, cli.List("v1/ticket-types", params, &resp)
}
