// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryBareMetal2Instance queries BareMetal2Instance list
func (cli *ZSClient) QueryBareMetal2Instance(params *param.QueryParam) ([]view.BareMetal2InstanceInventoryView, error) {
	var resp []view.BareMetal2InstanceInventoryView
	return resp, cli.List("v1/baremetal2/bm-instances", params, &resp)
}

func (cli *ZSClient) GetBareMetal2Instance(uuid string) (*view.BareMetal2InstanceInventoryView, error) {
	var resp view.BareMetal2InstanceInventoryView
	if err := cli.Get("v1/baremetal2/bm-instances", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// CreateBareMetal2Instance creates BareMetal2Instance
func (cli *ZSClient) CreateBareMetal2Instance(params param.CreateBareMetal2InstanceParam) (*view.BareMetal2InstanceInventoryView, error) {
	var resp view.CreateBareMetal2InstanceEventView
	if err := cli.Post("v1/baremetal2/bm-instances", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// CreateBareMetal2InstanceAsync Async
func (cli *ZSClient) CreateBareMetal2InstanceAsync(params param.CreateBareMetal2InstanceParam) (string, error) {

	resource := "v1/baremetal2/bm-instances"
	responseKey := ""
	var retVal interface{}

	apiId, err := cli.PostWithAsync(resource, responseKey, params, retVal, true)
	if err != nil {
		return "", err
	}

	return apiId, nil
}
// ReconnectBareMetal2Instance operates on BareMetal2Instance
func (cli *ZSClient) ReconnectBareMetal2Instance(uuid string, params param.ReconnectBareMetal2InstanceParam) (*view.BareMetal2InstanceInventoryView, error) {
	var resp view.ReconnectBareMetal2InstanceEventView
	if err := cli.Put("v1/baremetal2/bm-instances/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// StartBareMetal2Instance starts BareMetal2Instance
func (cli *ZSClient) StartBareMetal2Instance(uuid string, params param.StartBareMetal2InstanceParam) (*view.BareMetal2InstanceInventoryView, error) {
	var resp view.StartBareMetal2InstanceEventView
	if err := cli.Put("v1/baremetal2/bm-instances/{uuid}/action", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// UpdateBareMetal2Instance updates BareMetal2Instance
func (cli *ZSClient) UpdateBareMetal2Instance(uuid string, params param.UpdateBareMetal2InstanceParam) (*view.BareMetal2InstanceInventoryView, error) {
	var resp view.UpdateBareMetal2InstanceEventView
	if err := cli.Put("v1/baremetal2/bm-instances/{uuid}/action", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
