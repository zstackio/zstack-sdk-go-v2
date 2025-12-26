// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// DeleteBuildAppExportHistory deletes BuildAppExportHistory
func (cli *ZSClient) DeleteBuildAppExportHistory(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/buildapp/exports/{buildAppUuid}", uuid, string(deleteMode))
}
