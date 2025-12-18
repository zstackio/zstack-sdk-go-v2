// Copyright (c) ZStack.io, Inc.

package param

// AddBackendServerToServerGroupDetailParam AddBackendServerToServerGroup详细参数
type AddBackendServerToServerGroupDetailParam struct {
	rest string `json:"serverGroupUuid" validate:"required"` // 必填
	rest []interface{} `json:"vmNics,omitempty"`
	rest []interface{} `json:"servers,omitempty"`
}

// AddBackendServerToServerGroupParam AddBackendServerToServerGroup请求参数
type AddBackendServerToServerGroupParam struct {
	BaseParam
	Params AddBackendServerToServerGroupDetailParam `json:"params"` // 详细参数
}

