// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// CreateFlowCollector creates FlowCollector
func (cli *ZSClient) CreateFlowCollector(params param.CreateFlowCollectorParam) (*view.FlowCollectorInventoryView, error) {
	resp := view.FlowCollectorInventoryView{}
	if err := cli.Post("v1/flowmeters/{flowMeterUuid}/collectors", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// QueryFlowCollector queries FlowCollector list
func (cli *ZSClient) QueryFlowCollector(params *param.QueryParam) ([]view.FlowCollectorInventoryView, error) {
	var resp []view.FlowCollectorInventoryView
	return resp, cli.List("v1/flowmeters/collectors", params, &resp)
}

// PageFlowCollector Pagination
func (cli *ZSClient) PageFlowCollector(params *param.QueryParam) ([]view.FlowCollectorInventoryView, int, error) {
	var flowCollectors []view.FlowCollectorInventoryView
	total, err := cli.Page("v1/flowmeters/collectors", params, &flowCollectors)
	return flowCollectors, total, err
}
// UpdateFlowCollector updates FlowCollector
func (cli *ZSClient) UpdateFlowCollector(uuid string, params param.UpdateFlowCollectorParam) (*view.FlowCollectorInventoryView, error) {
	resp := view.FlowCollectorInventoryView{}
	if err := cli.Put("v1/flowmeters/collectors", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// DeleteFlowCollector deletes FlowCollector
func (cli *ZSClient) DeleteFlowCollector(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/flowmeters/collectors", uuid, string(deleteMode))
}
