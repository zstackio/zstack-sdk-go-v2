// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryIAM2VirtualIDAttribute queries IAM2VirtualIDAttribute list
func (cli *ZSClient) QueryIAM2VirtualIDAttribute(params param.QueryParam) ([]view.IAM2VirtualIDAttributeInventoryView, error) {
	var resp []view.IAM2VirtualIDAttributeInventoryView
	return resp, cli.List("v1/iam2/virtual-ids/attributes", &params, &resp)
}
