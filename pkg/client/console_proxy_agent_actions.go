// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// ReconnectConsoleProxyAgent operates on ConsoleProxyAgent
func (cli *ZSClient) ReconnectConsoleProxyAgent(uuid string, params param.ReconnectConsoleProxyAgentParam) (*view.ReconnectConsoleProxyAgentEventView, error) {
	resp := view.ReconnectConsoleProxyAgentEventView{}
	if err := cli.Put("v1/consoles/agents", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// QueryConsoleProxyAgent queries ConsoleProxyAgent list
func (cli *ZSClient) QueryConsoleProxyAgent(params *param.QueryParam) ([]view.ConsoleProxyAgentInventoryView, error) {
	var resp []view.ConsoleProxyAgentInventoryView
	return resp, cli.List("v1/consoles/agents", params, &resp)
}

// PageConsoleProxyAgent Pagination
func (cli *ZSClient) PageConsoleProxyAgent(params *param.QueryParam) ([]view.ConsoleProxyAgentInventoryView, int, error) {
	var consoleProxyAgents []view.ConsoleProxyAgentInventoryView
	total, err := cli.Page("v1/consoles/agents", params, &consoleProxyAgents)
	return consoleProxyAgents, total, err
}
// UpdateConsoleProxyAgent updates ConsoleProxyAgent
func (cli *ZSClient) UpdateConsoleProxyAgent(uuid string, params param.UpdateConsoleProxyAgentParam) (*view.ConsoleProxyAgentInventoryView, error) {
	resp := view.ConsoleProxyAgentInventoryView{}
	if err := cli.Put("v1/consoles/agents", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
