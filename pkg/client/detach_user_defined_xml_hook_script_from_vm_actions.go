// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// DetachUserDefinedXmlHookScriptFromVm operates on UserDefinedXmlHookScriptFromVm
func (cli *ZSClient) DetachUserDefinedXmlHookScriptFromVm(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/xmlhook/vm-instances/{vmInstanceUuid}/detach", uuid, string(deleteMode))
}
