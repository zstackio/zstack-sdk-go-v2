// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryEmailTriggerAction 查询EmailTriggerAction列表
func (cli *ZSClient) QueryEmailTriggerAction(params param.QueryParam) ([]view.QueryMonitorTriggerActionView, error) {
	var resp []view.QueryMonitorTriggerActionView
	return resp, cli.List("v1/monitoring/trigger-actions/emails", &params, &resp)
}

