// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// ReconnectBareMetal2Gateway operates on BareMetal2Gateway
func (cli *ZSClient) ReconnectBareMetal2Gateway(uuid string, params param.ReconnectBareMetal2GatewayParam) (*view.BareMetal2GatewayInventoryView, error) {
	var resp view.ReconnectBareMetal2GatewayEventView
	err := cli.PutWithSpec("v1/baremetal2/gateways", fmt.Sprintf(\"%s/actions\", uuid), params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// UpdateBareMetal2Gateway updates BareMetal2Gateway
func (cli *ZSClient) UpdateBareMetal2Gateway(uuid string, params param.UpdateBareMetal2GatewayParam) (*view.BareMetal2GatewayInventoryView, error) {
	var resp view.UpdateBareMetal2GatewayEventView
	err := cli.PutWithSpec("v1/baremetal2/gateways", fmt.Sprintf(\"%s/actions\", uuid), params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// QueryBareMetal2Gateway queries BareMetal2Gateway list
func (cli *ZSClient) QueryBareMetal2Gateway(params *param.QueryParam) ([]view.BareMetal2GatewayInventoryView, error) {
	var resp []view.BareMetal2GatewayInventoryView
	return resp, cli.List("v1/baremetal2/gateways", params, &resp)
}

func (cli *ZSClient) GetBareMetal2Gateway(uuid string) (*view.BareMetal2GatewayInventoryView, error) {
	var resp view.BareMetal2GatewayInventoryView
	if err := cli.Get("v1/baremetal2/gateways", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// DeleteBareMetal2Gateway deletes BareMetal2Gateway
func (cli *ZSClient) DeleteBareMetal2Gateway(uuid string, deleteMode param.DeleteMode) error {
	return cli.DeleteWithSpec("v1/baremetal2/gateways", fmt.Sprintf(\"%s\", uuid), string(deleteMode))
}
// AddBareMetal2Gateway adds BareMetal2Gateway
func (cli *ZSClient) AddBareMetal2Gateway(params param.AddBareMetal2GatewayParam) (*view.HostInventoryView, error) {
	var resp view.AddHostEventView
	if err := cli.Post("v1/baremetal2/gateways", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
