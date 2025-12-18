// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryPhysicalDriveSelfTestHistory 查询PhysicalDriveSelfTestHistory列表
func (cli *ZSClient) QueryPhysicalDriveSelfTestHistory(params param.QueryParam) ([]view.QueryPhysicalDriveSelfTestHistoryView, error) {
	var resp []view.QueryPhysicalDriveSelfTestHistoryView
	return resp, cli.List("v1/storage-devices/local-raid/physical-drives/self-test", &params, &resp)
}

