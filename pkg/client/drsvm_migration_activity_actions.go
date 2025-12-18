// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryDRSVmMigrationActivity 查询DRSVmMigrationActivity列表
func (cli *ZSClient) QueryDRSVmMigrationActivity(params param.QueryParam) ([]view.QueryDRSVmMigrationActivityView, error) {
	var resp []view.QueryDRSVmMigrationActivityView
	return resp, cli.List("v1/clusters/drs/vm-migration-activities", &params, &resp)
}

