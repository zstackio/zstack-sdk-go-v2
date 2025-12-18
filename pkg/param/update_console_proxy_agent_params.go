// Copyright (c) ZStack.io, Inc.

package param

// UpdateConsoleProxyAgentDetailParam UpdateConsoleProxyAgent detail param
type UpdateConsoleProxyAgentDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	ConsoleProxyOverriddenIp string `json:"consoleProxyOverriddenIp" validate:"required"`
	ConsoleProxyPort int `json:"consoleProxyPort,omitempty"`
}

// UpdateConsoleProxyAgentParam UpdateConsoleProxyAgent request param
type UpdateConsoleProxyAgentParam struct {
	BaseParam
	Params UpdateConsoleProxyAgentDetailParam `json:"params"`
}
