// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QueryHostOsCategory queries HostOsCategory list
func (cli *ZSClient) QueryHostOsCategory(params *param.QueryParam) ([]view.HostOsCategoryInventoryView, error) {
	var resp []view.HostOsCategoryInventoryView
	return resp, cli.List("v1/hosts/os/category", params, &resp)
}
