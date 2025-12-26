// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QueryIAM2Organization queries IAM2Organization list
func (cli *ZSClient) QueryIAM2Organization(params *param.QueryParam) ([]view.IAM2OrganizationInventoryView, error) {
	var resp []view.IAM2OrganizationInventoryView
	return resp, cli.List("v1/iam2/organizations", params, &resp)
}
