// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QueryConsoleProxyAgent queries ConsoleProxyAgent list
func (cli *ZSClient) QueryConsoleProxyAgent(params *param.QueryParam) ([]view.ConsoleProxyAgentInventoryView, error) {
	var resp []view.ConsoleProxyAgentInventoryView
	return resp, cli.List("v1/consoles/agents", params, &resp)
}
