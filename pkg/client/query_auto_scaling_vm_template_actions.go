// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QueryAutoScalingVmTemplate queries AutoScalingVmTemplate list
func (cli *ZSClient) QueryAutoScalingVmTemplate(params *param.QueryParam) ([]view.AutoScalingVmTemplateInventoryView, error) {
	var resp []view.AutoScalingVmTemplateInventoryView
	return resp, cli.List("v1/autoscaling/vmtemplate", params, &resp)
}
