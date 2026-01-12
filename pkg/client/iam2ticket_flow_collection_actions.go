// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// UpdateIAM2TicketFlowCollection updates IAM2TicketFlowCollection
func (cli *ZSClient) UpdateIAM2TicketFlowCollection(uuid string, params param.UpdateIAM2TicketFlowCollectionParam) (*view.TicketFlowCollectionInventoryView, error) {
	var resp view.UpdateTicketFlowCollectionEventView
	err := cli.PutWithSpec("v1/tickets/flow-collections", fmt.Sprintf(\"%s/actions\", uuid), params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
