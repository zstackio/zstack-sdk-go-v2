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
	return cli.Delete("v1/baremetal/pxeservers", uuid, string(deleteMode))
}
// UpdateBaremetalPxeServer updates BaremetalPxeServer
func (cli *ZSClient) UpdateBaremetalPxeServer(uuid string, params param.UpdateBaremetalPxeServerParam) (*view.BaremetalPxeServerInventoryView, error) {
	resp := view.BaremetalPxeServerInventoryView{}
	if err := cli.PutWithRespKey("v1/baremetal/pxeservers", uuid, "", map[string]interface{}{
		"updateBaremetalPxeServer": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// StartBaremetalPxeServer starts BaremetalPxeServer
func (cli *ZSClient) StartBaremetalPxeServer(uuid string, params param.StartBaremetalPxeServerParam) (*view.BaremetalPxeServerInventoryView, error) {
	resp := view.BaremetalPxeServerInventoryView{}
	if err := cli.PutWithRespKey("v1/baremetal/pxeservers", uuid, "", map[string]interface{}{
		"startBaremetalPxeServer": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// ReconnectBaremetalPxeServer operates on BaremetalPxeServer
func (cli *ZSClient) ReconnectBaremetalPxeServer(uuid string, params param.ReconnectBaremetalPxeServerParam) (*view.BaremetalPxeServerInventoryView, error) {
	resp := view.BaremetalPxeServerInventoryView{}
	if err := cli.PutWithRespKey("v1/baremetal/pxeservers", uuid, "", map[string]interface{}{
		"reconnectBaremetalPxeServer": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// StopBaremetalPxeServer stops BaremetalPxeServer
func (cli *ZSClient) StopBaremetalPxeServer(uuid string, params param.StopBaremetalPxeServerParam) (*view.BaremetalPxeServerInventoryView, error) {
	resp := view.BaremetalPxeServerInventoryView{}
	if err := cli.PutWithRespKey("v1/baremetal/pxeservers", uuid, "", map[string]interface{}{
		"stopBaremetalPxeServer": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
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

// PageBaremetalPxeServer Pagination
func (cli *ZSClient) PageBaremetalPxeServer(params *param.QueryParam) ([]view.BaremetalPxeServerInventoryView, int, error) {
	var baremetalPxeServers []view.BaremetalPxeServerInventoryView
	total, err := cli.Page("v1/baremetal/pxeservers", params, &baremetalPxeServers)
	return baremetalPxeServers, total, err
}
// CreateBaremetalPxeServer creates BaremetalPxeServer
func (cli *ZSClient) CreateBaremetalPxeServer(params param.CreateBaremetalPxeServerParam) (*view.BaremetalPxeServerInventoryView, error) {
	resp := view.BaremetalPxeServerInventoryView{}
	if err := cli.Post("v1/baremetal/pxeservers", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
