// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// DeleteVmBootMode 删除VmBootMode
func (cli *ZSClient) DeleteVmBootMode(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/vm-instances/{uuid}/bootmode", uuid, string(deleteMode))
}

