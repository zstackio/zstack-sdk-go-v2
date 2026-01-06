// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// CreateFlowCollector creates FlowCollector
func (cli *ZSClient) CreateFlowCollector(params param.CreateFlowCollectorParam) (*view.FlowCollectorInventoryView, error) {
	var resp view.CreateFlowCollectorEventView
	if err := cli.Post("v1/flowmeters/{flowMeterUuid}/collectors", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// QueryFlowCollector queries FlowCollector list
func (cli *ZSClient) QueryFlowCollector(params *param.QueryParam) ([]view.FlowCollectorInventoryView, error) {
	var resp []view.FlowCollectorInventoryView
	return resp, cli.List("v1/flowmeters/collectors", params, &resp)
}
// UpdateFlowCollector updates FlowCollector
func (cli *ZSClient) UpdateFlowCollector(uuid string, params param.UpdateFlowCollectorParam) (*view.FlowCollectorInventoryView, error) {
	var resp view.CreateFlowCollectorEventView
	if err := cli.Put("v1/flowmeters/collectors/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// DeleteFlowCollector deletes FlowCollector
func (cli *ZSClient) DeleteFlowCollector(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/flowmeters/collectors/{uuid}", uuid, string(deleteMode))
}
