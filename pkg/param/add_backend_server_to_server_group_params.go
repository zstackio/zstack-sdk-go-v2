// Copyright (c) ZStack.io, Inc.

package param

// AddBackendServerToServerGroupDetailParam AddBackendServerToServerGroup detail param
type AddBackendServerToServerGroupDetailParam struct {
	ServerGroupUuid string `json:"serverGroupUuid" validate:"required"`
	VmNics []interface{} `json:"vmNics,omitempty"`
	Servers []interface{} `json:"servers,omitempty"`
}

// AddBackendServerToServerGroupParam AddBackendServerToServerGroup request param
type AddBackendServerToServerGroupParam struct {
	BaseParam
	Params AddBackendServerToServerGroupDetailParam `json:"params"`
}
