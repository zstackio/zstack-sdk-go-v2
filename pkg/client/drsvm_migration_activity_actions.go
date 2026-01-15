// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryDRSVmMigrationActivity queries DRSVmMigrationActivity list
func (cli *ZSClient) QueryDRSVmMigrationActivity(params *param.QueryParam) ([]view.DRSVmMigrationActivityInventoryView, error) {
	var resp []view.DRSVmMigrationActivityInventoryView
	return resp, cli.List("v1/clusters/drs/vm-migration-activities", params, &resp)
}

// PageDRSVmMigrationActivity Pagination
func (cli *ZSClient) PageDRSVmMigrationActivity(params *param.QueryParam) ([]view.DRSVmMigrationActivityInventoryView, int, error) {
	var dRSVmMigrationActivities []view.DRSVmMigrationActivityInventoryView
	total, err := cli.Page("v1/clusters/drs/vm-migration-activities", params, &dRSVmMigrationActivities)
	return dRSVmMigrationActivities, total, err
}
