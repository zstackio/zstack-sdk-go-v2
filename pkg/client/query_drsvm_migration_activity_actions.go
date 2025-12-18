// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryDRSVmMigrationActivity queries DRSVmMigrationActivity list
func (cli *ZSClient) QueryDRSVmMigrationActivity(params param.QueryParam) ([]view.DRSVmMigrationActivityInventoryView, error) {
	var resp []view.DRSVmMigrationActivityInventoryView
	return resp, cli.List("v1/clusters/drs/vm-migration-activities", &params, &resp)
}
