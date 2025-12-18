// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// ExpungeImageGroup operates on ImageGroup
func (cli *ZSClient) ExpungeImageGroup(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/imagegroups/{uuid}/actions", uuid, string(deleteMode))
}
