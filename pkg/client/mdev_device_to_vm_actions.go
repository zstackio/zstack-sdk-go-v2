// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// AttachMdevDeviceToVm 操作MdevDeviceToVm
func (cli *ZSClient) AttachMdevDeviceToVm(params param.AttachMdevDeviceToVmParam) (*view.AttachMdevDeviceToVmEventView, error) {
	resp := view.AttachMdevDeviceToVmEventView{}
	if err := cli.Post("v1/mdev-devices/{mdevDeviceUuid}/vm-instances/{vmInstanceUuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

