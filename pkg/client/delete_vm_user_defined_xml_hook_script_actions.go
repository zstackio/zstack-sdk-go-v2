// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// DeleteVmUserDefinedXmlHookScript deletes VmUserDefinedXmlHookScript
func (cli *ZSClient) DeleteVmUserDefinedXmlHookScript(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/vm-instances/{vmInstanceUuid}/xml-hook-script", uuid, string(deleteMode))
}
