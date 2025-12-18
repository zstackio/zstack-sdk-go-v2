// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryIAM2ProjectAttribute queries IAM2ProjectAttribute list
func (cli *ZSClient) QueryIAM2ProjectAttribute(params param.QueryParam) ([]view.IAM2ProjectAttributeInventoryView, error) {
	var resp []view.IAM2ProjectAttributeInventoryView
	return resp, cli.List("v1/iam2/projects/attributes", &params, &resp)
}
