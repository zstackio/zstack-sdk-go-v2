// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryIAM2OrganizationAttribute 查询IAM2OrganizationAttribute列表
func (cli *ZSClient) QueryIAM2OrganizationAttribute(params param.QueryParam) ([]view.QueryIAM2OrganizationAttributeView, error) {
	var resp []view.QueryIAM2OrganizationAttributeView
	return resp, cli.List("v1/iam2/organizations/attributes", &params, &resp)
}

