// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// DeleteVmStaticIp deletes VmStaticIp
func (cli *ZSClient) DeleteVmStaticIp(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/vm-instances/{vmInstanceUuid}/static-ips", uuid, string(deleteMode))
}
