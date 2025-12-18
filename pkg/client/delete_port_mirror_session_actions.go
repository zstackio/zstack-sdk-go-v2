// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// DeletePortMirrorSession deletes PortMirrorSession
func (cli *ZSClient) DeletePortMirrorSession(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/port-mirrors/sessons/{uuid}", uuid, string(deleteMode))
}
