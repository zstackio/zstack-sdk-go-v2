// Copyright (c) ZStack.io, Inc.

package param

// ChangeAccessControlListServerGroupDetailParam ChangeAccessControlListServerGroup详细参数
type ChangeAccessControlListServerGroupDetailParam struct {
	rest []string `json:"serverGroupUuids" validate:"required"` // 必填
	rest string `json:"listenerUuid" validate:"required"` // 必填
	rest string `json:"aclUuid" validate:"required"` // 必填
}

// ChangeAccessControlListServerGroupParam ChangeAccessControlListServerGroup请求参数
type ChangeAccessControlListServerGroupParam struct {
	BaseParam
	Params ChangeAccessControlListServerGroupDetailParam `json:"params"` // 详细参数
}

