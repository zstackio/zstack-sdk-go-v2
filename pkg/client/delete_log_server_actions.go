// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// DeleteLogServer deletes LogServer
func (cli *ZSClient) DeleteLogServer(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/log/servers", uuid, string(deleteMode))
}
