// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryLicenseAuthorizedNode queries LicenseAuthorizedNode list
func (cli *ZSClient) QueryLicenseAuthorizedNode(params param.QueryParam) ([]view.LicenseAuthorizedNodeInventoryView, error) {
	var resp []view.LicenseAuthorizedNodeInventoryView
	return resp, cli.List("v1/license-servers", &params, &resp)
}
