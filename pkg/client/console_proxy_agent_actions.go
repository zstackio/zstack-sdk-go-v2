// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryConsoleProxyAgent 查询ConsoleProxyAgent列表
func (cli *ZSClient) QueryConsoleProxyAgent(params param.QueryParam) ([]view.QueryConsoleProxyAgentView, error) {
	var resp []view.QueryConsoleProxyAgentView
	return resp, cli.List("v1/consoles/agents", &params, &resp)
}

