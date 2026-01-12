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

func (cli *ZSClient) GetConsoleProxyAgent(uuid string) (*view.ConsoleProxyAgentInventoryView, error) {
	var resp view.ConsoleProxyAgentInventoryView
	if err := cli.Get("v1/consoles/agents", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// UpdateConsoleProxyAgent updates ConsoleProxyAgent
func (cli *ZSClient) UpdateConsoleProxyAgent(uuid string, params param.UpdateConsoleProxyAgentParam) (*view.ConsoleProxyAgentInventoryView, error) {
	var resp view.UpdateConsoleProxyAgentEventView
	err := cli.PutWithSpec("v1/consoles/agents", fmt.Sprintf(\"%s/actions\", uuid), params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
