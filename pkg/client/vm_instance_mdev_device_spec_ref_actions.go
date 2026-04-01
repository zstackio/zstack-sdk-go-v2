// Copyright (c) ZStack.io, Inc.

package client

import (
	"fmt"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryVmInstanceMdevDeviceSpecRef queries VmInstanceMdevDeviceSpecRef list
func (cli *ZSClient) QueryVmInstanceMdevDeviceSpecRef(vmInstanceUuid string, params *param.QueryParam) ([]view.VmInstanceMdevDeviceSpecRefInventoryView, error) {
	var resp []view.VmInstanceMdevDeviceSpecRefInventoryView
	return resp, cli.List(fmt.Sprintf("v1/vm-instances/%s/mdev-device-specs", vmInstanceUuid), params, &resp)
}

func (cli *ZSClient) GetVmInstanceMdevDeviceSpecRef(uuid string) (*view.VmInstanceMdevDeviceSpecRefInventoryView, error) {
	var resp view.VmInstanceMdevDeviceSpecRefInventoryView
	if err := cli.Get("v1/vm-instances", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageVmInstanceMdevDeviceSpecRef Pagination
func (cli *ZSClient) PageVmInstanceMdevDeviceSpecRef(vmInstanceUuid string, params *param.QueryParam) ([]view.VmInstanceMdevDeviceSpecRefInventoryView, int, error) {
	var vmInstanceMdevDeviceSpecRefs []view.VmInstanceMdevDeviceSpecRefInventoryView
	total, err := cli.Page(fmt.Sprintf("v1/vm-instances/%s/mdev-device-specs", vmInstanceUuid), params, &vmInstanceMdevDeviceSpecRefs)
	return vmInstanceMdevDeviceSpecRefs, total, err
}
