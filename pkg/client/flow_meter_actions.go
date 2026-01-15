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

// PageFlowMeter Pagination
func (cli *ZSClient) PageFlowMeter(params *param.QueryParam) ([]view.FlowMeterInventoryView, int, error) {
	var flowMeters []view.FlowMeterInventoryView
	total, err := cli.Page("v1/flowmeters", params, &flowMeters)
	return flowMeters, total, err
}
// CreateFlowMeter creates FlowMeter
func (cli *ZSClient) CreateFlowMeter(params param.CreateFlowMeterParam) (*view.FlowMeterInventoryView, error) {
	resp := view.FlowMeterInventoryView{}
	if err := cli.Post("v1/flowmeters", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// DeleteFlowMeter deletes FlowMeter
func (cli *ZSClient) DeleteFlowMeter(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/flowmeters", uuid, string(deleteMode))
}
// UpdateFlowMeter updates FlowMeter
func (cli *ZSClient) UpdateFlowMeter(uuid string, params param.UpdateFlowMeterParam) (*view.FlowMeterInventoryView, error) {
	resp := view.FlowMeterInventoryView{}
	if err := cli.Put("v1/flowmeters", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
