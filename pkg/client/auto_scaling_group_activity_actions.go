// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryAutoScalingGroupActivity 查询AutoScalingGroupActivity列表
func (cli *ZSClient) QueryAutoScalingGroupActivity(params param.QueryParam) ([]view.QueryAutoScalingGroupActivityView, error) {
	var resp []view.QueryAutoScalingGroupActivityView
	return resp, cli.List("v1/autoscaling/groups/activities", &params, &resp)
}

