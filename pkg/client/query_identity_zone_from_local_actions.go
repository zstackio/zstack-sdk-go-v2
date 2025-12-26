// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QueryIdentityZoneFromLocal queries IdentityZoneFromLocal list
func (cli *ZSClient) QueryIdentityZoneFromLocal(params *param.QueryParam) ([]view.IdentityZoneInventoryView, error) {
	var resp []view.IdentityZoneInventoryView
	return resp, cli.List("v1/hybrid/identity-zone", params, &resp)
}
