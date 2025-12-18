// Copyright (c) ZStack.io, Inc.

package param

// AddRolesToIAM2VirtualIDGroupDetailParam AddRolesToIAM2VirtualIDGroup详细参数
type AddRolesToIAM2VirtualIDGroupDetailParam struct {
	rest []string `json:"roleUuids" validate:"required"` // 必填
	rest string `json:"groupUuid" validate:"required"` // 必填
	rest string `json:"projectUuid,omitempty"`
}

// AddRolesToIAM2VirtualIDGroupParam AddRolesToIAM2VirtualIDGroup请求参数
type AddRolesToIAM2VirtualIDGroupParam struct {
	BaseParam
	Params AddRolesToIAM2VirtualIDGroupDetailParam `json:"params"` // 详细参数
}

