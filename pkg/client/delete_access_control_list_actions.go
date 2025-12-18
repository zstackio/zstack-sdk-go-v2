// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// DeleteAccessControlList deletes AccessControlList
func (cli *ZSClient) DeleteAccessControlList(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/access-control-lists/{uuid}", uuid, string(deleteMode))
}
