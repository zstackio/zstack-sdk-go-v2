// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryLicenseAuthorizedNode queries LicenseAuthorizedNode list
func (cli *ZSClient) QueryLicenseAuthorizedNode(params *param.QueryParam) ([]view.LicenseAuthorizedNodeInventoryView, error) {
	var resp []view.LicenseAuthorizedNodeInventoryView
	return resp, cli.List("v1/license-servers", params, &resp)
}
