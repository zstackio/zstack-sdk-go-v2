// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// RemovePciDeviceSpecFromVmInstance 操作RemovePciDeviceSpecFromVmInstance
func (cli *ZSClient) RemovePciDeviceSpecFromVmInstance(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/pci-device-specs/{pciSpecUuid}/vm-instances/{vmInstanceUuid}", uuid, string(deleteMode))
}

