// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// ExpungeVmUserDefinedXmlHookScript operates on VmUserDefinedXmlHookScript
func (cli *ZSClient) ExpungeVmUserDefinedXmlHookScript(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/vm-instances/xml-hook-script/{uuid}", uuid, string(deleteMode))
}
