// Copyright (c) ZStack.io, Inc.

package client

import (
	"context"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryLicenseAuthorizedNode queries LicenseAuthorizedNode list
func (cli *ZSClient) QueryLicenseAuthorizedNode(ctx context.Context, params *param.QueryParam) ([]view.LicenseAuthorizedNodeInventoryView, error) {
	var resp []view.LicenseAuthorizedNodeInventoryView
	return resp, cli.List(ctx, "v1/license-servers", params, &resp)
}

func (cli *ZSClient) GetLicenseAuthorizedNode(ctx context.Context, uuid string) (*view.LicenseAuthorizedNodeInventoryView, error) {
	var resp view.LicenseAuthorizedNodeInventoryView
	if err := cli.Get(ctx, "v1/license-servers", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageLicenseAuthorizedNode Pagination
func (cli *ZSClient) PageLicenseAuthorizedNode(ctx context.Context, params *param.QueryParam) ([]view.LicenseAuthorizedNodeInventoryView, int, error) {
	var licenseAuthorizedNodes []view.LicenseAuthorizedNodeInventoryView
	total, err := cli.Page(ctx, "v1/license-servers", params, &licenseAuthorizedNodes)
	return licenseAuthorizedNodes, total, err
}
