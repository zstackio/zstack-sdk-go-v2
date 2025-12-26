// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// DeleteIscsiServer deletes IscsiServer
func (cli *ZSClient) DeleteIscsiServer(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/storage-devices/iscsi/servers/{uuid}", uuid, string(deleteMode))
}
