// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryIAM2VirtualIDGroupAttribute queries IAM2VirtualIDGroupAttribute list
func (cli *ZSClient) QueryIAM2VirtualIDGroupAttribute(params param.QueryParam) ([]view.IAM2VirtualIDGroupAttributeInventoryView, error) {
	var resp []view.IAM2VirtualIDGroupAttributeInventoryView
	return resp, cli.List("v1/iam2/projects/groups/attributes/", &params, &resp)
}
