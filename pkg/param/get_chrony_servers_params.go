// Copyright (c) ZStack.io, Inc.

package param

// GetChronyServersDetailParam GetChronyServers detail param
type GetChronyServersDetailParam struct {
}

// GetChronyServersParam GetChronyServers request param
type GetChronyServersParam struct {
	BaseParam
	Params GetChronyServersDetailParam `json:"params"`
}
