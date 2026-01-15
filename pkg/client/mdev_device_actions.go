// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// UpdateMdevDevice updates MdevDevice
func (cli *ZSClient) UpdateMdevDevice(uuid string, params param.UpdateMdevDeviceParam) (*view.MdevDeviceInventoryView, error) {
	resp := view.MdevDeviceInventoryView{}
	if err := cli.Put("v1/mdev-devices", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// DeleteMdevDevice deletes MdevDevice
func (cli *ZSClient) DeleteMdevDevice(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/mdev-devices", uuid, string(deleteMode))
}
// QueryMdevDevice queries MdevDevice list
func (cli *ZSClient) QueryMdevDevice(params *param.QueryParam) ([]view.MdevDeviceInventoryView, error) {
	var resp []view.MdevDeviceInventoryView
	return resp, cli.List("v1/mdev-devices", params, &resp)
}

// PageMdevDevice Pagination
func (cli *ZSClient) PageMdevDevice(params *param.QueryParam) ([]view.MdevDeviceInventoryView, int, error) {
	var mdevDevices []view.MdevDeviceInventoryView
	total, err := cli.Page("v1/mdev-devices", params, &mdevDevices)
	return mdevDevices, total, err
}
