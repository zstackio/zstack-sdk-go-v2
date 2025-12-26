// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QueryIAM2VirtualIDAttribute queries IAM2VirtualIDAttribute list
func (cli *ZSClient) QueryIAM2VirtualIDAttribute(params *param.QueryParam) ([]view.IAM2VirtualIDAttributeInventoryView, error) {
	var resp []view.IAM2VirtualIDAttributeInventoryView
	return resp, cli.List("v1/iam2/virtual-ids/attributes", params, &resp)
}
