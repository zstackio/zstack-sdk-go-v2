// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// ReconnectConsoleProxyAgent operates on ConsoleProxyAgent
func (cli *ZSClient) ReconnectConsoleProxyAgent(ctx context.Context, params param.ReconnectConsoleProxyAgentParam) (*view.ReconnectConsoleProxyAgentEventView, error) {
	resp := view.ReconnectConsoleProxyAgentEventView{}
	if err := cli.PutWithRespKey(ctx, "v1/consoles/agents", "", "", map[string]interface{}{
		"reconnectConsoleProxyAgent": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// QueryConsoleProxyAgent queries ConsoleProxyAgent list
func (cli *ZSClient) QueryConsoleProxyAgent(ctx context.Context, params *param.QueryParam) ([]view.ConsoleProxyAgentInventoryView, error) {
	var resp []view.ConsoleProxyAgentInventoryView
	return resp, cli.List(ctx, "v1/consoles/agents", params, &resp)
}

func (cli *ZSClient) GetConsoleProxyAgent(ctx context.Context, uuid string) (*view.ConsoleProxyAgentInventoryView, error) {
	var resp view.ConsoleProxyAgentInventoryView
	if err := cli.Get(ctx, "v1/consoles/agents", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageConsoleProxyAgent Pagination
func (cli *ZSClient) PageConsoleProxyAgent(ctx context.Context, params *param.QueryParam) ([]view.ConsoleProxyAgentInventoryView, int, error) {
	var consoleProxyAgents []view.ConsoleProxyAgentInventoryView
	total, err := cli.Page(ctx, "v1/consoles/agents", params, &consoleProxyAgents)
	return consoleProxyAgents, total, err
}
// UpdateConsoleProxyAgent updates ConsoleProxyAgent
func (cli *ZSClient) UpdateConsoleProxyAgent(ctx context.Context, uuid string, params param.UpdateConsoleProxyAgentParam) (*view.ConsoleProxyAgentInventoryView, error) {
	resp := view.ConsoleProxyAgentInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/consoles/agents", uuid, "", map[string]interface{}{
		"updateConsoleProxyAgent": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
