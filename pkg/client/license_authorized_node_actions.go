// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryLicenseAuthorizedNode queries LicenseAuthorizedNode list
func (cli *ZSClient) QueryLicenseAuthorizedNode(params *param.QueryParam) ([]view.LicenseAuthorizedNodeInventoryView, error) {
	var resp []view.LicenseAuthorizedNodeInventoryView
	return resp, cli.List("v1/license-servers", params, &resp)
}

// PageLicenseAuthorizedNode Pagination
func (cli *ZSClient) PageLicenseAuthorizedNode(params *param.QueryParam) ([]view.LicenseAuthorizedNodeInventoryView, int, error) {
	var licenseAuthorizedNodes []view.LicenseAuthorizedNodeInventoryView
	total, err := cli.Page("v1/license-servers", params, &licenseAuthorizedNodes)
	return licenseAuthorizedNodes, total, err
}
