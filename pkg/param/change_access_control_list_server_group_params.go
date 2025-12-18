// Copyright (c) ZStack.io, Inc.

package param

// ChangeAccessControlListServerGroupDetailParam ChangeAccessControlListServerGroup detail param
type ChangeAccessControlListServerGroupDetailParam struct {
	ServerGroupUuids []string `json:"serverGroupUuids" validate:"required"`
	ListenerUuid string `json:"listenerUuid" validate:"required"`
	AclUuid string `json:"aclUuid" validate:"required"`
}

// ChangeAccessControlListServerGroupParam ChangeAccessControlListServerGroup request param
type ChangeAccessControlListServerGroupParam struct {
	BaseParam
	Params ChangeAccessControlListServerGroupDetailParam `json:"params"`
}
