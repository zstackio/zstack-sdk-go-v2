// Copyright (c) ZStack.io, Inc.

package param

// ReconnectConsoleProxyAgentDetailParam ReconnectConsoleProxyAgent detail param
type ReconnectConsoleProxyAgentDetailParam struct {
	AgentUuids []string `json:"agentUuids,omitempty"`
}

// ReconnectConsoleProxyAgentParam ReconnectConsoleProxyAgent request param
type ReconnectConsoleProxyAgentParam struct {
	BaseParam
	Params ReconnectConsoleProxyAgentDetailParam `json:"params"`
}
