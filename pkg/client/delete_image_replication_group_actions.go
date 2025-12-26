// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// DeleteImageReplicationGroup deletes ImageReplicationGroup
func (cli *ZSClient) DeleteImageReplicationGroup(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/image-replication-groups/{uuid}", uuid, string(deleteMode))
}
