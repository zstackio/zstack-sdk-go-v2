// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// DetachUserDefinedXmlHookScriptFromVm operates on UserDefinedXmlHookScriptFromVm
func (cli *ZSClient) DetachUserDefinedXmlHookScriptFromVm(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/xmlhook/vm-instances/{vmInstanceUuid}/detach", uuid, string(deleteMode))
}
