// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QueryEventFromResourceStack queries EventFromResourceStack list
func (cli *ZSClient) QueryEventFromResourceStack(params *param.QueryParam) ([]view.CloudFormationStackEventInventoryView, error) {
	var resp []view.CloudFormationStackEventInventoryView
	return resp, cli.List("v1/cloudformation/event", params, &resp)
}
