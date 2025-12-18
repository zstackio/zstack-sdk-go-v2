// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryIdentityZoneFromLocal queries IdentityZoneFromLocal list
func (cli *ZSClient) QueryIdentityZoneFromLocal(params param.QueryParam) ([]view.IdentityZoneInventoryView, error) {
	var resp []view.IdentityZoneInventoryView
	return resp, cli.List("v1/hybrid/identity-zone", &params, &resp)
}
