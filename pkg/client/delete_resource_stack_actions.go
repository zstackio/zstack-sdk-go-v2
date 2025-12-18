// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// DeleteResourceStack deletes ResourceStack
func (cli *ZSClient) DeleteResourceStack(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/cloudformation/stack/{uuid}", uuid, string(deleteMode))
}
