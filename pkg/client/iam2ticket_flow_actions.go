// Copyright (c) ZStack.io, Inc.

package client

import (
	"context"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// UpdateIAM2TicketFlow updates IAM2TicketFlow
func (cli *ZSClient) UpdateIAM2TicketFlow(ctx context.Context, uuid string, params param.UpdateIAM2TicketFlowParam) (*view.TicketFlowInventoryView, error) {
	resp := view.TicketFlowInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/tickets/flow", uuid, "", map[string]interface{}{
		"updateIAM2TicketFlow": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// DeleteIAM2TicketFlow deletes IAM2TicketFlow
func (cli *ZSClient) DeleteIAM2TicketFlow(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/tickets/flow", uuid, string(deleteMode))
}
// AddIAM2TicketFlow adds IAM2TicketFlow
func (cli *ZSClient) AddIAM2TicketFlow(ctx context.Context, params param.AddIAM2TicketFlowParam) (*view.TicketFlowInventoryView, error) {
	resp := view.TicketFlowInventoryView{}
	if err := cli.PostWithRespKey(ctx, "v1/tickets/flow", "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
