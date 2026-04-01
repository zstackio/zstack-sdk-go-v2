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
	resp := view.SdnControllerInventoryView{}
	if err := cli.Post("v1/sdn-controllers", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// UpdateSdnController updates SdnController
func (cli *ZSClient) UpdateSdnController(uuid string, params param.UpdateSdnControllerParam) (*view.SdnControllerInventoryView, error) {
	resp := view.SdnControllerInventoryView{}
	if err := cli.PutWithRespKey("v1/sdn-controllers", uuid, "", map[string]interface{}{
		"updateSdnController": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// ChangeSdnController changes SdnController
func (cli *ZSClient) ChangeSdnController(uuid string, params param.ChangeSdnControllerParam) (*view.SdnControllerInventoryView, error) {
	resp := view.SdnControllerInventoryView{}
	if err := cli.PutWithRespKey("v1/sdn-controllers", uuid, "", map[string]interface{}{
		"changeSdnController": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// ReconnectSdnController operates on SdnController
func (cli *ZSClient) ReconnectSdnController(sdnControllerUuid string, params param.ReconnectSdnControllerParam) (*view.SdnControllerInventoryView, error) {
	resp := view.SdnControllerInventoryView{}
	if err := cli.PutWithRespKey("v1/sdn-controllers", sdnControllerUuid, "", map[string]interface{}{
		"reconnectSdnController": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
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

// PageSdnController Pagination
func (cli *ZSClient) PageSdnController(params *param.QueryParam) ([]view.SdnControllerInventoryView, int, error) {
	var sdnControllers []view.SdnControllerInventoryView
	total, err := cli.Page("v1/sdn-controllers", params, &sdnControllers)
	return sdnControllers, total, err
}
