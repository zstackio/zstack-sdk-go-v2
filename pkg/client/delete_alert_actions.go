// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// DeleteAlert deletes Alert
func (cli *ZSClient) DeleteAlert(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/monitoring/alerts", uuid, string(deleteMode))
}
