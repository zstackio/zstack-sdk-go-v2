// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QuerySharedResource 查询SharedResource列表
func (cli *ZSClient) QuerySharedResource(params param.QueryParam) ([]view.QuerySharedResourceView, error) {
	var resp []view.QuerySharedResourceView
	return resp, cli.List("v1/accounts/resources", &params, &resp)
}

