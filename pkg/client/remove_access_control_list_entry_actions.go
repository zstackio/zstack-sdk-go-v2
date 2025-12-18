// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// RemoveAccessControlListEntry 操作RemoveAccessControlListEntry
func (cli *ZSClient) RemoveAccessControlListEntry(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/access-control-lists/{aclUuid}/ipentries/{uuid}", uuid, string(deleteMode))
}

