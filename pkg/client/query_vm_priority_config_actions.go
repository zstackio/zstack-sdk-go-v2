// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryVmPriorityConfig queries VmPriorityConfig list
func (cli *ZSClient) QueryVmPriorityConfig(params param.QueryParam) ([]view.VmPriorityConfigInventoryView, error) {
	var resp []view.VmPriorityConfigInventoryView
	return resp, cli.List("v1/vm-priority-config", &params, &resp)
}
