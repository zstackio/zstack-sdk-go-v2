// Copyright (c) ZStack.io, Inc.

package param

// RemoveRolesFromIAM2VirtualIDGroupDetailParam RemoveRolesFromIAM2VirtualIDGroup详细参数
type RemoveRolesFromIAM2VirtualIDGroupDetailParam struct {
	rest []string `json:"roleUuids" validate:"required"` // 必填
	rest string `json:"groupUuid" validate:"required"` // 必填
	rest string `json:"projectUuid,omitempty"`
}

// RemoveRolesFromIAM2VirtualIDGroupParam RemoveRolesFromIAM2VirtualIDGroup请求参数
type RemoveRolesFromIAM2VirtualIDGroupParam struct {
	BaseParam
	Params RemoveRolesFromIAM2VirtualIDGroupDetailParam `json:"params"` // 详细参数
}

