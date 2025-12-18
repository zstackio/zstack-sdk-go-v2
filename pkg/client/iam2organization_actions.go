// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryIAM2Organization 查询IAM2Organization列表
func (cli *ZSClient) QueryIAM2Organization(params param.QueryParam) ([]view.QueryIAM2OrganizationView, error) {
	var resp []view.QueryIAM2OrganizationView
	return resp, cli.List("v1/iam2/organizations", &params, &resp)
}

