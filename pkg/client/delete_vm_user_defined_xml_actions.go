// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// DeleteVmUserDefinedXml deletes VmUserDefinedXml
func (cli *ZSClient) DeleteVmUserDefinedXml(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/vm-instances/{vmInstanceUuid}/xml", uuid, string(deleteMode))
}
