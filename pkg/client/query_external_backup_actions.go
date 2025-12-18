// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryExternalBackup queries ExternalBackup list
func (cli *ZSClient) QueryExternalBackup(params param.QueryParam) ([]view.ExternalBackupInventoryView, error) {
	var resp []view.ExternalBackupInventoryView
	return resp, cli.List("v1/externalbackup", &params, &resp)
}
