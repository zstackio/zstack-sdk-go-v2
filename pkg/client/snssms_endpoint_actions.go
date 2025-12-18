// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QuerySNSSmsEndpoint 查询SNSSmsEndpoint列表
func (cli *ZSClient) QuerySNSSmsEndpoint(params param.QueryParam) ([]view.QuerySNSSmsEndpointView, error) {
	var resp []view.QuerySNSSmsEndpointView
	return resp, cli.List("v1/sns/sms-endpoints", &params, &resp)
}

