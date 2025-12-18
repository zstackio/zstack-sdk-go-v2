// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryIAM2OrganizationProjectRef 查询IAM2OrganizationProjectRef列表
func (cli *ZSClient) QueryIAM2OrganizationProjectRef(params param.QueryParam) ([]view.QueryIAM2OrganizationProjectRefView, error) {
	var resp []view.QueryIAM2OrganizationProjectRefView
	return resp, cli.List("v1/iam2/projects/organizations/refs", &params, &resp)
}

