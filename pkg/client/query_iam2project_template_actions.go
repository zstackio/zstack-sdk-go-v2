// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QueryIAM2ProjectTemplate queries IAM2ProjectTemplate list
func (cli *ZSClient) QueryIAM2ProjectTemplate(params *param.QueryParam) ([]view.IAM2ProjectTemplateInventoryView, error) {
	var resp []view.IAM2ProjectTemplateInventoryView
	return resp, cli.List("v1/iam2/projects/templates", params, &resp)
}
