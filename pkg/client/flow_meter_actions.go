// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryFlowMeter queries FlowMeter list
func (cli *ZSClient) QueryFlowMeter(params *param.QueryParam) ([]view.FlowMeterInventoryView, error) {
	var resp []view.FlowMeterInventoryView
	return resp, cli.List("v1/flowmeters", params, &resp)
}

func (cli *ZSClient) GetFlowMeter(uuid string) (*view.FlowMeterInventoryView, error) {
	var resp view.FlowMeterInventoryView
	if err := cli.Get("v1/flowmeters", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
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
	return cli.Delete("v1/flowmeters", uuid, string(deleteMode))
}
// UpdateFlowMeter updates FlowMeter
func (cli *ZSClient) UpdateFlowMeter(uuid string, params param.UpdateFlowMeterParam) (*view.FlowMeterInventoryView, error) {
	var resp view.UpdateFlowMeterEventView
	if err := cli.Put("v1/flowmeters", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
