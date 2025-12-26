// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// RemoveAccessControlListEntry removes AccessControlListEntry
func (cli *ZSClient) RemoveAccessControlListEntry(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/access-control-lists/{aclUuid}/ipentries/{uuid}", uuid, string(deleteMode))
}
