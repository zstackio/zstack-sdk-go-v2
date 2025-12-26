// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// ExpungeVmUserDefinedXmlHookScript operates on VmUserDefinedXmlHookScript
func (cli *ZSClient) ExpungeVmUserDefinedXmlHookScript(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/vm-instances/xml-hook-script/{uuid}", uuid, string(deleteMode))
}
