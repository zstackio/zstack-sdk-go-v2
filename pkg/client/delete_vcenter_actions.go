// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// DeleteVCenter deletes VCenter
func (cli *ZSClient) DeleteVCenter(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/vcenters/{uuid}", uuid, string(deleteMode))
}
