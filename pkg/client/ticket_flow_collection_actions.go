// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryTicketFlowCollection queries TicketFlowCollection list
func (cli *ZSClient) QueryTicketFlowCollection(params *param.QueryParam) ([]view.TicketFlowCollectionInventoryView, error) {
	var resp []view.TicketFlowCollectionInventoryView
	return resp, cli.List("v1/tickets/flow-collections", params, &resp)
}

func (cli *ZSClient) GetTicketFlowCollection(uuid string) (*view.TicketFlowCollectionInventoryView, error) {
	var resp view.TicketFlowCollectionInventoryView
	if err := cli.Get("v1/tickets/flow-collections", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// DeleteTicketFlowCollection deletes TicketFlowCollection
func (cli *ZSClient) DeleteTicketFlowCollection(uuid string, deleteMode param.DeleteMode) error {
	return cli.DeleteWithSpec("v1/tickets/flow-collections", fmt.Sprintf(\"%s\", uuid), string(deleteMode))
}
