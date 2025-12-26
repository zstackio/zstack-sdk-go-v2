// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// ReconnectConsoleProxyAgent operates on ReconnectConsoleProxyAgent
func (cli *ZSClient) ReconnectConsoleProxyAgent(uuid string, params param.ReconnectConsoleProxyAgentParam) (*view.ReconnectConsoleProxyAgentEventView, error) {
	resp := view.ReconnectConsoleProxyAgentEventView{}
	if err := cli.Put("v1/consoles/agents", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
