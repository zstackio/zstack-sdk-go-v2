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

func (cli *ZSClient) GetLicenseAuthorizedNode(uuid string) (*view.LicenseAuthorizedNodeInventoryView, error) {
	var resp view.LicenseAuthorizedNodeInventoryView
	if err := cli.Get("v1/license-servers", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
