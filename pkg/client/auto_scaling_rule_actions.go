// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryAutoScalingRule 查询AutoScalingRule列表
func (cli *ZSClient) QueryAutoScalingRule(params param.QueryParam) ([]view.QueryAutoScalingRuleView, error) {
	var resp []view.QueryAutoScalingRuleView
	return resp, cli.List("v1/autoscaling/groups/rules", &params, &resp)
}

