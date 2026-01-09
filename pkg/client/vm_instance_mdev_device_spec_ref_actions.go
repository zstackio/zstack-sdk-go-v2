// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryVmInstanceMdevDeviceSpecRef queries VmInstanceMdevDeviceSpecRef list
func (cli *ZSClient) QueryVmInstanceMdevDeviceSpecRef(params *param.QueryParam) ([]view.VmInstanceMdevDeviceSpecRefInventoryView, error) {
	var resp []view.VmInstanceMdevDeviceSpecRefInventoryView
	return resp, cli.List("v1/vm-instances/{vmInstanceUuid}/mdev-device-specs", params, &resp)
}

func (cli *ZSClient) GetVmInstanceMdevDeviceSpecRef(uuid string) (*view.VmInstanceMdevDeviceSpecRefInventoryView, error) {
	var resp view.VmInstanceMdevDeviceSpecRefInventoryView
	if err := cli.Get("v1/vm-instances/{vmInstanceUuid}/mdev-device-specs", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
