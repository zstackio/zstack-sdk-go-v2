// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QueryIAM2VirtualIDGroup queries IAM2VirtualIDGroup list
func (cli *ZSClient) QueryIAM2VirtualIDGroup(params *param.QueryParam) ([]view.IAM2VirtualIDGroupInventoryView, error) {
	var resp []view.IAM2VirtualIDGroupInventoryView
	return resp, cli.List("v1/iam2/projects/groups", params, &resp)
}
