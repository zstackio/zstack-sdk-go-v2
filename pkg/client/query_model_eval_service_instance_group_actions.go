// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryModelEvalServiceInstanceGroup queries ModelEvalServiceInstanceGroup list
func (cli *ZSClient) QueryModelEvalServiceInstanceGroup(params param.QueryParam) ([]view.ModelServiceInstanceGroupInventoryView, error) {
	var resp []view.ModelServiceInstanceGroupInventoryView
	return resp, cli.List("v1/ai/model-eval-services/instances/groups/", &params, &resp)
}
