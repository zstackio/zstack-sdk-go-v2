// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QuerySNSDingTalkEndpoint 查询SNSDingTalkEndpoint列表
func (cli *ZSClient) QuerySNSDingTalkEndpoint(params param.QueryParam) ([]view.QuerySNSDingTalkEndpointView, error) {
	var resp []view.QuerySNSDingTalkEndpointView
	return resp, cli.List("v1/sns/application-endpoints/ding-talk", &params, &resp)
}

