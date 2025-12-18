// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryIAM2ProjectAccountRef queries IAM2ProjectAccountRef list
func (cli *ZSClient) QueryIAM2ProjectAccountRef(params param.QueryParam) ([]view.IAM2ProjectAccountRefInventoryView, error) {
	var resp []view.IAM2ProjectAccountRefInventoryView
	return resp, cli.List("v1/iam2/projects/account/refs", &params, &resp)
}
