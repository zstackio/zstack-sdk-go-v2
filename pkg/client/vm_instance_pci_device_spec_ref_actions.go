// Copyright (c) ZStack.io, Inc.

package client

import (
	"fmt"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryVmInstancePciDeviceSpecRef queries VmInstancePciDeviceSpecRef list
func (cli *ZSClient) QueryVmInstancePciDeviceSpecRef(vmInstanceUuid string, params *param.QueryParam) ([]view.VmInstancePciDeviceSpecRefInventoryView, error) {
	var resp []view.VmInstancePciDeviceSpecRefInventoryView
	return resp, cli.List(fmt.Sprintf("v1/vm-instances/%s/pci-device-specs", vmInstanceUuid), params, &resp)
}

func (cli *ZSClient) GetVmInstancePciDeviceSpecRef(uuid string) (*view.VmInstancePciDeviceSpecRefInventoryView, error) {
	var resp view.VmInstancePciDeviceSpecRefInventoryView
	if err := cli.Get("v1/vm-instances", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageVmInstancePciDeviceSpecRef Pagination
func (cli *ZSClient) PageVmInstancePciDeviceSpecRef(vmInstanceUuid string, params *param.QueryParam) ([]view.VmInstancePciDeviceSpecRefInventoryView, int, error) {
	var vmInstancePciDeviceSpecRefs []view.VmInstancePciDeviceSpecRefInventoryView
	total, err := cli.Page(fmt.Sprintf("v1/vm-instances/%s/pci-device-specs", vmInstanceUuid), params, &vmInstancePciDeviceSpecRefs)
	return vmInstancePciDeviceSpecRefs, total, err
}
