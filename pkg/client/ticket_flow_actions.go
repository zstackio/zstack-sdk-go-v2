// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryTicketFlow queries TicketFlow list
func (cli *ZSClient) QueryTicketFlow(params *param.QueryParam) ([]view.TicketFlowInventoryView, error) {
	var resp []view.TicketFlowInventoryView
	return resp, cli.List("v1/tickets/flow", params, &resp)
}

func (cli *ZSClient) GetTicketFlow(uuid string) (*view.TicketFlowInventoryView, error) {
	var resp view.TicketFlowInventoryView
	if err := cli.Get("v1/tickets/flow", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
