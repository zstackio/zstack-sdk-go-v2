// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryIAM2Project queries IAM2Project list
func (cli *ZSClient) QueryIAM2Project(params param.QueryParam) ([]view.IAM2ProjectInventoryView, error) {
	var resp []view.IAM2ProjectInventoryView
	return resp, cli.List("v1/iam2/projects", &params, &resp)
}
