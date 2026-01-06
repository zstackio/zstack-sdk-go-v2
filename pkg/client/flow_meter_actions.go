// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryFlowMeter queries FlowMeter list
func (cli *ZSClient) QueryFlowMeter(params *param.QueryParam) ([]view.FlowMeterInventoryView, error) {
	var resp []view.FlowMeterInventoryView
	return resp, cli.List("v1/flowmeters", params, &resp)
}
// CreateFlowMeter creates FlowMeter
func (cli *ZSClient) CreateFlowMeter(params param.CreateFlowMeterParam) (*view.FlowMeterInventoryView, error) {
	var resp view.CreateFlowMeterEventView
	if err := cli.Post("v1/flowmeters", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// DeleteFlowMeter deletes FlowMeter
func (cli *ZSClient) DeleteFlowMeter(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/flowmeters/{uuid}", uuid, string(deleteMode))
}
// UpdateFlowMeter updates FlowMeter
func (cli *ZSClient) UpdateFlowMeter(uuid string, params param.UpdateFlowMeterParam) (*view.FlowMeterInventoryView, error) {
	var resp view.UpdateFlowMeterEventView
	if err := cli.Put("v1/flowmeters/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
