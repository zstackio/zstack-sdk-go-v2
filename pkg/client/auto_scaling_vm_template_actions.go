// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryAutoScalingVmTemplate 查询AutoScalingVmTemplate列表
func (cli *ZSClient) QueryAutoScalingVmTemplate(params param.QueryParam) ([]view.QueryAutoScalingVmTemplateView, error) {
	var resp []view.QueryAutoScalingVmTemplateView
	return resp, cli.List("v1/autoscaling/vmtemplate", &params, &resp)
}

