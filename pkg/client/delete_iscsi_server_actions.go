// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// DeleteIscsiServer deletes IscsiServer
func (cli *ZSClient) DeleteIscsiServer(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/storage-devices/iscsi/servers/{uuid}", uuid, string(deleteMode))
}
