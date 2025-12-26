// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QueryZBoxBackup queries ZBoxBackup list
func (cli *ZSClient) QueryZBoxBackup(params *param.QueryParam) ([]view.ZBoxBackupInventoryView, error) {
	var resp []view.ZBoxBackupInventoryView
	return resp, cli.List("v1/externalbackup/zbox", params, &resp)
}
