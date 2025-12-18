// Copyright (c) ZStack.io, Inc.

package param

// RemoveBackendServerFromServerGroupDetailParam RemoveBackendServerFromServerGroup detail param
type RemoveBackendServerFromServerGroupDetailParam struct {
	ServerGroupUuid string `json:"serverGroupUuid" validate:"required"`
	VmNicUuids []string `json:"vmNicUuids,omitempty"`
	ServerIps []string `json:"serverIps,omitempty"`
}

// RemoveBackendServerFromServerGroupParam RemoveBackendServerFromServerGroup request param
type RemoveBackendServerFromServerGroupParam struct {
	BaseParam
	Params RemoveBackendServerFromServerGroupDetailParam `json:"params"`
}
