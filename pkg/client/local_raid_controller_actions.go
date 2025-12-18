// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryLocalRaidController 查询LocalRaidController列表
func (cli *ZSClient) QueryLocalRaidController(params param.QueryParam) ([]view.QueryLocalRaidPhysicalDriveView, error) {
	var resp []view.QueryLocalRaidPhysicalDriveView
	return resp, cli.List("v1/storage-devices/local-raid/controllers", &params, &resp)
}

