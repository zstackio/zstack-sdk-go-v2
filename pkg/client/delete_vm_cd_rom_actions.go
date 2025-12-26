// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// DeleteVmCdRom deletes VmCdRom
func (cli *ZSClient) DeleteVmCdRom(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/vm-instances/cdroms/{uuid}", uuid, string(deleteMode))
}
