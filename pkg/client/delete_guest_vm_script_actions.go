// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// DeleteGuestVmScript deletes GuestVmScript
func (cli *ZSClient) DeleteGuestVmScript(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/scripts/{uuid}", uuid, string(deleteMode))
}
