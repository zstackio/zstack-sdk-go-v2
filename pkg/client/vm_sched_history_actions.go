// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryVmSchedHistory 查询VmSchedHistory列表
func (cli *ZSClient) QueryVmSchedHistory(params param.QueryParam) ([]view.QueryVmSchedHistoryView, error) {
	var resp []view.QueryVmSchedHistoryView
	return resp, cli.List("v1/vm/sched-history", &params, &resp)
}

