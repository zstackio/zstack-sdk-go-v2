// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QueryIAM2VirtualIDGroupAttribute queries IAM2VirtualIDGroupAttribute list
func (cli *ZSClient) QueryIAM2VirtualIDGroupAttribute(params *param.QueryParam) ([]view.IAM2VirtualIDGroupAttributeInventoryView, error) {
	var resp []view.IAM2VirtualIDGroupAttributeInventoryView
	return resp, cli.List("v1/iam2/projects/groups/attributes/", params, &resp)
}
