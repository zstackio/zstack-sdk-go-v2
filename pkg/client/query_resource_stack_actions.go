// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QueryResourceStack queries ResourceStack list
func (cli *ZSClient) QueryResourceStack(params *param.QueryParam) ([]view.ResourceStackInventoryView, error) {
	var resp []view.ResourceStackInventoryView
	return resp, cli.List("v1/cloudformation/stack", params, &resp)
}
