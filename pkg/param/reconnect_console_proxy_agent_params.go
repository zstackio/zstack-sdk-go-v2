// Copyright (c) ZStack.io, Inc.

package param

// ReconnectConsoleProxyAgentDetailParam ReconnectConsoleProxyAgent详细参数
type ReconnectConsoleProxyAgentDetailParam struct {
	rest []string `json:"agentUuids,omitempty"`
}

// ReconnectConsoleProxyAgentParam ReconnectConsoleProxyAgent请求参数
type ReconnectConsoleProxyAgentParam struct {
	BaseParam
	Params ReconnectConsoleProxyAgentDetailParam `json:"params"` // 详细参数
}

