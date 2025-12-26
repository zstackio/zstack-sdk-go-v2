// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QueryIAM2ProjectRole queries IAM2ProjectRole list
func (cli *ZSClient) QueryIAM2ProjectRole(params *param.QueryParam) ([]view.IAM2ProjectRoleInventoryView, error) {
	var resp []view.IAM2ProjectRoleInventoryView
	return resp, cli.List("v1/iam2/project-roles", params, &resp)
}
