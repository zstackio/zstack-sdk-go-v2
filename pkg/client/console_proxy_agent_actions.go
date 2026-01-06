// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
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
// UpdateConsoleProxyAgent updates ConsoleProxyAgent
func (cli *ZSClient) UpdateConsoleProxyAgent(uuid string, params param.UpdateConsoleProxyAgentParam) (*view.ConsoleProxyAgentInventoryView, error) {
	var resp view.UpdateConsoleProxyAgentEventView
	if err := cli.Put("v1/consoles/agents/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
