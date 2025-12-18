// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryEventLog 查询EventLog列表
func (cli *ZSClient) QueryEventLog(params param.QueryParam) ([]view.QueryEventLogView, error) {
	var resp []view.QueryEventLogView
	return resp, cli.List("v1/eventlogs", &params, &resp)
}

