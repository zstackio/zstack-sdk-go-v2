// Copyright (c) ZStack.io, Inc.

package param

// AddRolesToIAM2VirtualIDDetailParam AddRolesToIAM2VirtualID详细参数
type AddRolesToIAM2VirtualIDDetailParam struct {
	rest string `json:"virtualIDUuid" validate:"required"` // 必填
	rest []string `json:"roleUuids" validate:"required"` // 必填
	rest string `json:"projectUuid,omitempty"`
}

// AddRolesToIAM2VirtualIDParam AddRolesToIAM2VirtualID请求参数
type AddRolesToIAM2VirtualIDParam struct {
	BaseParam
	Params AddRolesToIAM2VirtualIDDetailParam `json:"params"` // 详细参数
}

