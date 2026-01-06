// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// UpdateMdevDevice updates MdevDevice
func (cli *ZSClient) UpdateMdevDevice(uuid string, params param.UpdateMdevDeviceParam) (*view.MdevDeviceInventoryView, error) {
	var resp view.UpdateMdevDeviceEventView
	if err := cli.Put("v1/mdev-devices/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// DeleteMdevDevice deletes MdevDevice
func (cli *ZSClient) DeleteMdevDevice(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/mdev-devices/{mdevDeviceUuid}", uuid, string(deleteMode))
}
// QueryMdevDevice queries MdevDevice list
func (cli *ZSClient) QueryMdevDevice(params *param.QueryParam) ([]view.MdevDeviceInventoryView, error) {
	var resp []view.MdevDeviceInventoryView
	return resp, cli.List("v1/mdev-devices", params, &resp)
}
