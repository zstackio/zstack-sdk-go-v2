// Copyright (c) ZStack.io, Inc.

package param

// RemoveRolesFromIAM2VirtualIDDetailParam RemoveRolesFromIAM2VirtualID详细参数
type RemoveRolesFromIAM2VirtualIDDetailParam struct {
	rest []string `json:"roleUuids" validate:"required"` // 必填
	rest string `json:"virtualIDUuid" validate:"required"` // 必填
	rest string `json:"projectUuid,omitempty"`
}

// RemoveRolesFromIAM2VirtualIDParam RemoveRolesFromIAM2VirtualID请求参数
type RemoveRolesFromIAM2VirtualIDParam struct {
	BaseParam
	Params RemoveRolesFromIAM2VirtualIDDetailParam `json:"params"` // 详细参数
}

