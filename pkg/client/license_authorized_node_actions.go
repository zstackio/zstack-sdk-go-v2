// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryLicenseAuthorizedNode 查询LicenseAuthorizedNode列表
func (cli *ZSClient) QueryLicenseAuthorizedNode(params param.QueryParam) ([]view.QueryLicenseAuthorizedNodeView, error) {
	var resp []view.QueryLicenseAuthorizedNodeView
	return resp, cli.List("v1/license-servers", &params, &resp)
}

