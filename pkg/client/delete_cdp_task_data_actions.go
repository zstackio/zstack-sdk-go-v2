// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// DeleteCdpTaskData deletes CdpTaskData
func (cli *ZSClient) DeleteCdpTaskData(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/cdp-task/{uuid}/data", uuid, string(deleteMode))
}
