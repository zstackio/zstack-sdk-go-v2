// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// DeleteCdpTask deletes CdpTask
func (cli *ZSClient) DeleteCdpTask(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/cdp-task/{uuid}", uuid, string(deleteMode))
}
