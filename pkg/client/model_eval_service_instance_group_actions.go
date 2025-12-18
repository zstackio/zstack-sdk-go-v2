// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryModelEvalServiceInstanceGroup 查询ModelEvalServiceInstanceGroup列表
func (cli *ZSClient) QueryModelEvalServiceInstanceGroup(params param.QueryParam) ([]view.QueryModelServiceInstanceGroupView, error) {
	var resp []view.QueryModelServiceInstanceGroupView
	return resp, cli.List("v1/ai/model-eval-services/instances/groups/", &params, &resp)
}

