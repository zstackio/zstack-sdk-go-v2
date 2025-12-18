// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// DeleteBaremetalPxeServer deletes BaremetalPxeServer
func (cli *ZSClient) DeleteBaremetalPxeServer(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/baremetal/pxeservers/{uuid}", uuid, string(deleteMode))
}
