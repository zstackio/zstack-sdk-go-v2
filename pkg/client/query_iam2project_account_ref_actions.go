// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QueryIAM2ProjectAccountRef queries IAM2ProjectAccountRef list
func (cli *ZSClient) QueryIAM2ProjectAccountRef(params *param.QueryParam) ([]view.IAM2ProjectAccountRefInventoryView, error) {
	var resp []view.IAM2ProjectAccountRefInventoryView
	return resp, cli.List("v1/iam2/projects/account/refs", params, &resp)
}
