// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// DetachMdevDeviceFromVm 操作MdevDeviceFromVm
func (cli *ZSClient) DetachMdevDeviceFromVm(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/mdev-devices/{mdevDeviceUuid}/vm-instances/{vmInstanceUuid}", uuid, string(deleteMode))
}

