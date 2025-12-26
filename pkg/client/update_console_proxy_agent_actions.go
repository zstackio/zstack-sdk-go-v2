// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// UpdateConsoleProxyAgent updates ConsoleProxyAgent
func (cli *ZSClient) UpdateConsoleProxyAgent(uuid string, params param.UpdateConsoleProxyAgentParam) (*view.UpdateConsoleProxyAgentEventView, error) {
	resp := view.UpdateConsoleProxyAgentEventView{}
	if err := cli.Put("v1/consoles/agents/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
