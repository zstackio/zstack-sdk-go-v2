// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// DeleteBaremetalPxeServer deletes BaremetalPxeServer
func (cli *ZSClient) DeleteBaremetalPxeServer(uuid string, deleteMode param.DeleteMode) error {
	return cli.DeleteWithSpec("v1/baremetal/pxeservers", fmt.Sprintf(\"%s\", uuid), string(deleteMode))
}
// UpdateBaremetalPxeServer updates BaremetalPxeServer
func (cli *ZSClient) UpdateBaremetalPxeServer(uuid string, params param.UpdateBaremetalPxeServerParam) (*view.BaremetalPxeServerInventoryView, error) {
	var resp view.UpdateBaremetalPxeServerEventView
	err := cli.PutWithSpec("v1/baremetal/pxeservers", fmt.Sprintf(\"%s/actions\", uuid), params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// StartBaremetalPxeServer starts BaremetalPxeServer
func (cli *ZSClient) StartBaremetalPxeServer(uuid string, params param.StartBaremetalPxeServerParam) (*view.BaremetalPxeServerInventoryView, error) {
	var resp view.StartBaremetalPxeServerEventView
	err := cli.PutWithSpec("v1/baremetal/pxeservers", fmt.Sprintf(\"%s/actions\", uuid), params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// ReconnectBaremetalPxeServer operates on BaremetalPxeServer
func (cli *ZSClient) ReconnectBaremetalPxeServer(uuid string, params param.ReconnectBaremetalPxeServerParam) (*view.BaremetalPxeServerInventoryView, error) {
	var resp view.ReconnectBaremetalPxeServerEventView
	err := cli.PutWithSpec("v1/baremetal/pxeservers", fmt.Sprintf(\"%s/actions\", uuid), params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// StopBaremetalPxeServer stops BaremetalPxeServer
func (cli *ZSClient) StopBaremetalPxeServer(uuid string, params param.StopBaremetalPxeServerParam) (*view.BaremetalPxeServerInventoryView, error) {
	var resp view.StopBaremetalPxeServerEventView
	err := cli.PutWithSpec("v1/baremetal/pxeservers", fmt.Sprintf(\"%s/actions\", uuid), params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// QueryBaremetalPxeServer queries BaremetalPxeServer list
func (cli *ZSClient) QueryBaremetalPxeServer(params *param.QueryParam) ([]view.BaremetalPxeServerInventoryView, error) {
	var resp []view.BaremetalPxeServerInventoryView
	return resp, cli.List("v1/baremetal/pxeservers", params, &resp)
}

func (cli *ZSClient) GetBaremetalPxeServer(uuid string) (*view.BaremetalPxeServerInventoryView, error) {
	var resp view.BaremetalPxeServerInventoryView
	if err := cli.Get("v1/baremetal/pxeservers", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// CreateBaremetalPxeServer creates BaremetalPxeServer
func (cli *ZSClient) CreateBaremetalPxeServer(params param.CreateBaremetalPxeServerParam) (*view.BaremetalPxeServerInventoryView, error) {
	var resp view.CreateBaremetalPxeServerEventView
	if err := cli.Post("v1/baremetal/pxeservers", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
