// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// RemoveMdevDeviceSpecFromVmInstance 操作RemoveMdevDeviceSpecFromVmInstance
func (cli *ZSClient) RemoveMdevDeviceSpecFromVmInstance(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/mdev-device-specs/{mdevSpecUuid}/vm-instances/{vmInstanceUuid}", uuid, string(deleteMode))
}

