// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryAgentVersion 查询AgentVersion列表
func (cli *ZSClient) QueryAgentVersion(params param.QueryParam) ([]view.QueryAgentVersionView, error) {
	var resp []view.QueryAgentVersionView
	return resp, cli.List("v1/agent-version", &params, &resp)
}

