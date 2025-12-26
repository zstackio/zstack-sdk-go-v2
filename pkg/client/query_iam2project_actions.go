// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QueryIAM2Project queries IAM2Project list
func (cli *ZSClient) QueryIAM2Project(params *param.QueryParam) ([]view.IAM2ProjectInventoryView, error) {
	var resp []view.IAM2ProjectInventoryView
	return resp, cli.List("v1/iam2/projects", params, &resp)
}
