// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// DeleteVmStaticIp 删除VmStaticIp
func (cli *ZSClient) DeleteVmStaticIp(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/vm-instances/{vmInstanceUuid}/static-ips", uuid, string(deleteMode))
}

