// Copyright (c) ZStack.io, Inc.

package param

// UpdateChronyServersDetailParam UpdateChronyServers detail param
type UpdateChronyServersDetailParam struct {
	InternalHostnames []string `json:"internalHostnames,omitempty"`
	ExternalHostnames []string `json:"externalHostnames,omitempty"`
}

// UpdateChronyServersParam UpdateChronyServers request param
type UpdateChronyServersParam struct {
	BaseParam
	Params UpdateChronyServersDetailParam `json:"params"`
}
