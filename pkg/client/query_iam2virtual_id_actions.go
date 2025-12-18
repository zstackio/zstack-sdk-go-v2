// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryIAM2VirtualID queries IAM2VirtualID list
func (cli *ZSClient) QueryIAM2VirtualID(params param.QueryParam) ([]view.IAM2VirtualIDInventoryView, error) {
	var resp []view.IAM2VirtualIDInventoryView
	return resp, cli.List("v1/iam2/virtual-ids", &params, &resp)
}
