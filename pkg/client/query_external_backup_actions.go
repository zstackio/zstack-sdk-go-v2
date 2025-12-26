// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QueryExternalBackup queries ExternalBackup list
func (cli *ZSClient) QueryExternalBackup(params *param.QueryParam) ([]view.ExternalBackupInventoryView, error) {
	var resp []view.ExternalBackupInventoryView
	return resp, cli.List("v1/externalbackup", params, &resp)
}
