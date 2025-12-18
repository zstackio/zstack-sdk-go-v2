// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryVmSchedulingRule 查询VmSchedulingRule列表
func (cli *ZSClient) QueryVmSchedulingRule(params param.QueryParam) ([]view.QueryVmSchedulingRuleView, error) {
	var resp []view.QueryVmSchedulingRuleView
	return resp, cli.List("v1/query/vm/schedulingRule", &params, &resp)
}

