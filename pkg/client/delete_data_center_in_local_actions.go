// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// DeleteDataCenterInLocal deletes DataCenterInLocal
func (cli *ZSClient) DeleteDataCenterInLocal(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/hybrid/data-center/{uuid}", uuid, string(deleteMode))
}
