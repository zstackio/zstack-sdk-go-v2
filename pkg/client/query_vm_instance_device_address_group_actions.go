// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryVmInstanceDeviceAddressGroup queries VmInstanceDeviceAddressGroup list
func (cli *ZSClient) QueryVmInstanceDeviceAddressGroup(params param.QueryParam) ([]view.VmInstanceDeviceAddressGroupInventoryView, error) {
	var resp []view.VmInstanceDeviceAddressGroupInventoryView
	return resp, cli.List("v1/vmInstance/device/address/group", &params, &resp)
}
