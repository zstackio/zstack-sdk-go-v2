// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QueryIAM2OrganizationAttribute queries IAM2OrganizationAttribute list
func (cli *ZSClient) QueryIAM2OrganizationAttribute(params *param.QueryParam) ([]view.IAM2OrganizationAttributeInventoryView, error) {
	var resp []view.IAM2OrganizationAttributeInventoryView
	return resp, cli.List("v1/iam2/organizations/attributes", params, &resp)
}
