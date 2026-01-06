// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// DeleteBaremetalPxeServer deletes BaremetalPxeServer
func (cli *ZSClient) DeleteBaremetalPxeServer(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/baremetal/pxeservers/{uuid}", uuid, string(deleteMode))
}
// UpdateBaremetalPxeServer updates BaremetalPxeServer
func (cli *ZSClient) UpdateBaremetalPxeServer(uuid string, params param.UpdateBaremetalPxeServerParam) (*view.BaremetalPxeServerInventoryView, error) {
	var resp view.UpdateBaremetalPxeServerEventView
	if err := cli.Put("v1/baremetal/pxeservers/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// StartBaremetalPxeServer starts BaremetalPxeServer
func (cli *ZSClient) StartBaremetalPxeServer(uuid string, params param.StartBaremetalPxeServerParam) (*view.BaremetalPxeServerInventoryView, error) {
	var resp view.StartBaremetalPxeServerEventView
	if err := cli.Put("v1/baremetal/pxeservers/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// ReconnectBaremetalPxeServer operates on BaremetalPxeServer
func (cli *ZSClient) ReconnectBaremetalPxeServer(uuid string, params param.ReconnectBaremetalPxeServerParam) (*view.BaremetalPxeServerInventoryView, error) {
	var resp view.ReconnectBaremetalPxeServerEventView
	if err := cli.Put("v1/baremetal/pxeservers/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// StopBaremetalPxeServer stops BaremetalPxeServer
func (cli *ZSClient) StopBaremetalPxeServer(uuid string, params param.StopBaremetalPxeServerParam) (*view.BaremetalPxeServerInventoryView, error) {
	var resp view.StopBaremetalPxeServerEventView
	if err := cli.Put("v1/baremetal/pxeservers/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// QueryBaremetalPxeServer queries BaremetalPxeServer list
func (cli *ZSClient) QueryBaremetalPxeServer(params *param.QueryParam) ([]view.BaremetalPxeServerInventoryView, error) {
	var resp []view.BaremetalPxeServerInventoryView
	return resp, cli.List("v1/baremetal/pxeservers", params, &resp)
}
// CreateBaremetalPxeServer creates BaremetalPxeServer
func (cli *ZSClient) CreateBaremetalPxeServer(params param.CreateBaremetalPxeServerParam) (*view.BaremetalPxeServerInventoryView, error) {
	var resp view.CreateBaremetalPxeServerEventView
	if err := cli.Post("v1/baremetal/pxeservers", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
