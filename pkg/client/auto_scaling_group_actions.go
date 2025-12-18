// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryAutoScalingGroup 查询AutoScalingGroup列表
func (cli *ZSClient) QueryAutoScalingGroup(params param.QueryParam) ([]view.QueryAutoScalingGroupView, error) {
	var resp []view.QueryAutoScalingGroupView
	return resp, cli.List("v1/autoscaling/groups", &params, &resp)
}

