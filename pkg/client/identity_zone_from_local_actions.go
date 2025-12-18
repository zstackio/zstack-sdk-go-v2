// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryIdentityZoneFromLocal 查询IdentityZoneFromLocal列表
func (cli *ZSClient) QueryIdentityZoneFromLocal(params param.QueryParam) ([]view.QueryIdentityZoneFromLocalView, error) {
	var resp []view.QueryIdentityZoneFromLocalView
	return resp, cli.List("v1/hybrid/identity-zone", &params, &resp)
}

