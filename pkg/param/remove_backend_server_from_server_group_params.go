// Copyright (c) ZStack.io, Inc.

package param

// RemoveBackendServerFromServerGroupDetailParam RemoveBackendServerFromServerGroup详细参数
type RemoveBackendServerFromServerGroupDetailParam struct {
	rest string `json:"serverGroupUuid" validate:"required"` // 必填
	rest []string `json:"vmNicUuids,omitempty"`
	rest []string `json:"serverIps,omitempty"`
}

// RemoveBackendServerFromServerGroupParam RemoveBackendServerFromServerGroup请求参数
type RemoveBackendServerFromServerGroupParam struct {
	BaseParam
	Params RemoveBackendServerFromServerGroupDetailParam `json:"params"` // 详细参数
}

