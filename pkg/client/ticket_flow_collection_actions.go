// Copyright (c) ZStack.io, Inc.

package client

import (
	"fmt"
	"net/url"

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

// PageTicketFlowCollection Pagination
func (cli *ZSClient) PageTicketFlowCollection(params *param.QueryParam) ([]view.TicketFlowCollectionInventoryView, int, error) {
	var ticketFlowCollections []view.TicketFlowCollectionInventoryView
	total, err := cli.Page("v1/tickets/flow-collections", params, &ticketFlowCollections)
	return ticketFlowCollections, total, err
}
// DeleteTicketFlowCollection deletes TicketFlowCollection
func (cli *ZSClient) DeleteTicketFlowCollection(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/tickets/flow-collections", uuid, string(deleteMode))
}

// RemoveTicketTypesFromTicketFlowCollection requires additional arguments.
//
// Deprecated: use RemoveTicketTypesFromTicketFlowCollectionWithParams to supply the required fields.
func (cli *ZSClient) RemoveTicketTypesFromTicketFlowCollection(uuid string, deleteMode param.DeleteMode) error {
	return fmt.Errorf("RemoveTicketTypesFromTicketFlowCollection requires additional arguments; use RemoveTicketTypesFromTicketFlowCollectionWithParams")
}

// RemoveTicketTypesFromTicketFlowCollectionWithParams sends the DELETE request with its operation parameters.
func (cli *ZSClient) RemoveTicketTypesFromTicketFlowCollectionWithParams(ticketFlowCollectionUuid string, deleteMode param.DeleteMode, params param.RemoveTicketTypesFromTicketFlowCollectionParamDetail) error {
	query := url.Values{
		"ticketTypeUuids": params.TicketTypeUuids,
	}
	return cli.deleteActionWithQuery("v1/tickets/flow-collections", ticketFlowCollectionUuid, "ticket-types", deleteMode, query, "ticketTypeUuids")
}

// ChangeTicketFlowCollectionState changes TicketFlowCollectionState
func (cli *ZSClient) ChangeTicketFlowCollectionState(uuid string, params param.ChangeTicketFlowCollectionStateParam) (*view.TicketFlowCollectionInventoryView, error) {
	resp := view.TicketFlowCollectionInventoryView{}
	if err := cli.PutWithRespKey("v1/tickets/flow-collections", uuid, "", map[string]interface{}{
		"changeTicketFlowCollectionState": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AddTicketTypesToTicketFlowCollection adds TicketTypesToTicketFlowCollection
func (cli *ZSClient) AddTicketTypesToTicketFlowCollection(ticketFlowCollectionUuid string, params param.AddTicketTypesToTicketFlowCollectionParam) (*view.AddTicketTypesToTicketFlowCollectionEventView, error) {
	resp := view.AddTicketTypesToTicketFlowCollectionEventView{}
	if err := cli.Post(fmt.Sprintf("v1/tickets/flow-collections/%s/ticket-types", ticketFlowCollectionUuid), params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
