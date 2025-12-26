// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QueryVmInstanceMdevDeviceSpecRef queries VmInstanceMdevDeviceSpecRef list
func (cli *ZSClient) QueryVmInstanceMdevDeviceSpecRef(params *param.QueryParam) ([]view.VmInstanceMdevDeviceSpecRefInventoryView, error) {
	var resp []view.VmInstanceMdevDeviceSpecRefInventoryView
	return resp, cli.List("v1/vm-instances/{vmInstanceUuid}/mdev-device-specs", params, &resp)
}
