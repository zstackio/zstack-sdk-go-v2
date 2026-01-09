// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryTicketType queries TicketType list
func (cli *ZSClient) QueryTicketType(params *param.QueryParam) ([]view.TicketTypeInventoryView, error) {
	var resp []view.TicketTypeInventoryView
	return resp, cli.List("v1/ticket-types", params, &resp)
}

func (cli *ZSClient) GetTicketType(uuid string) (*view.TicketTypeInventoryView, error) {
	var resp view.TicketTypeInventoryView
	if err := cli.Get("v1/ticket-types", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
