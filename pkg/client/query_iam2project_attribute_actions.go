// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QueryIAM2ProjectAttribute queries IAM2ProjectAttribute list
func (cli *ZSClient) QueryIAM2ProjectAttribute(params *param.QueryParam) ([]view.IAM2ProjectAttributeInventoryView, error) {
	var resp []view.IAM2ProjectAttributeInventoryView
	return resp, cli.List("v1/iam2/projects/attributes", params, &resp)
}
