// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QueryIAM2VirtualID queries IAM2VirtualID list
func (cli *ZSClient) QueryIAM2VirtualID(params *param.QueryParam) ([]view.IAM2VirtualIDInventoryView, error) {
	var resp []view.IAM2VirtualIDInventoryView
	return resp, cli.List("v1/iam2/virtual-ids", params, &resp)
}
