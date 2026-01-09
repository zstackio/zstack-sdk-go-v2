// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryVmInstanceDeviceAddressGroup queries VmInstanceDeviceAddressGroup list
func (cli *ZSClient) QueryVmInstanceDeviceAddressGroup(params *param.QueryParam) ([]view.VmInstanceDeviceAddressGroupInventoryView, error) {
	var resp []view.VmInstanceDeviceAddressGroupInventoryView
	return resp, cli.List("v1/vmInstance/device/address/group", params, &resp)
}

func (cli *ZSClient) GetVmInstanceDeviceAddressGroup(uuid string) (*view.VmInstanceDeviceAddressGroupInventoryView, error) {
	var resp view.VmInstanceDeviceAddressGroupInventoryView
	if err := cli.Get("v1/vmInstance/device/address/group", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
