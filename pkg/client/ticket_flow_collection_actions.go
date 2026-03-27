// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryTicketFlowCollection queries TicketFlowCollection list
func (cli *ZSClient) QueryTicketFlowCollection(ctx context.Context, params *param.QueryParam) ([]view.TicketFlowCollectionInventoryView, error) {
	var resp []view.TicketFlowCollectionInventoryView
	return resp, cli.List(ctx, "v1/tickets/flow-collections", params, &resp)
}

func (cli *ZSClient) GetTicketFlowCollection(ctx context.Context, uuid string) (*view.TicketFlowCollectionInventoryView, error) {
	var resp view.TicketFlowCollectionInventoryView
	if err := cli.Get(ctx, "v1/tickets/flow-collections", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageTicketFlowCollection Pagination
func (cli *ZSClient) PageTicketFlowCollection(ctx context.Context, params *param.QueryParam) ([]view.TicketFlowCollectionInventoryView, int, error) {
	var ticketFlowCollections []view.TicketFlowCollectionInventoryView
	total, err := cli.Page(ctx, "v1/tickets/flow-collections", params, &ticketFlowCollections)
	return ticketFlowCollections, total, err
}
// DeleteTicketFlowCollection deletes TicketFlowCollection
func (cli *ZSClient) DeleteTicketFlowCollection(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/tickets/flow-collections", uuid, string(deleteMode))
}
