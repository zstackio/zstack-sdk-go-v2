// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// RemoveSdnController removes SdnController
func (cli *ZSClient) RemoveSdnController(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/sdn-controllers", uuid, string(deleteMode))
}
// AddSdnController adds SdnController
func (cli *ZSClient) AddSdnController(params param.AddSdnControllerParam) (*view.SdnControllerInventoryView, error) {
	var resp view.AddSdnControllerEventView
	if err := cli.Post("v1/sdn-controllers", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// UpdateSdnController updates SdnController
func (cli *ZSClient) UpdateSdnController(uuid string, params param.UpdateSdnControllerParam) (*view.SdnControllerInventoryView, error) {
	var resp view.UpdateSdnControllerEventView
	if err := cli.Put("v1/sdn-controllers/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// ChangeSdnController changes SdnController
func (cli *ZSClient) ChangeSdnController(uuid string, params param.ChangeSdnControllerParam) (*view.SdnControllerInventoryView, error) {
	var resp view.ChangeSdnControllerEventView
	if err := cli.Put("v1/sdn-controllers/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// ReconnectSdnController operates on SdnController
func (cli *ZSClient) ReconnectSdnController(uuid string, params param.ReconnectSdnControllerParam) (*view.SdnControllerInventoryView, error) {
	var resp view.ReconnectSdnControllerEventView
	if err := cli.Put("v1/sdn-controllers/{sdnControllerUuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// QuerySdnController queries SdnController list
func (cli *ZSClient) QuerySdnController(params *param.QueryParam) ([]view.SdnControllerInventoryView, error) {
	var resp []view.SdnControllerInventoryView
	return resp, cli.List("v1/sdn-controllers", params, &resp)
}

func (cli *ZSClient) GetSdnController(uuid string) (*view.SdnControllerInventoryView, error) {
	var resp view.SdnControllerInventoryView
	if err := cli.Get("v1/sdn-controllers", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
