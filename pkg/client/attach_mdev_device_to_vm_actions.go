// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// AttachMdevDeviceToVm operates on MdevDeviceToVm
func (cli *ZSClient) AttachMdevDeviceToVm(params param.AttachMdevDeviceToVmParam) (*view.AttachMdevDeviceToVmEventView, error) {
	resp := view.AttachMdevDeviceToVmEventView{}
	if err := cli.Post("v1/mdev-devices/{mdevDeviceUuid}/vm-instances/{vmInstanceUuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
