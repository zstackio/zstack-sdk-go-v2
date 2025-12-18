// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// DeleteVmNic deletes VmNic
func (cli *ZSClient) DeleteVmNic(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/nics/{uuid}", uuid, string(deleteMode))
}
