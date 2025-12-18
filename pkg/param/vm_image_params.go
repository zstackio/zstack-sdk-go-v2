// Copyright (c) ZStack.io, Inc.

package param

// ChangeVmImageDetailParam ChangeVmImage详细参数
type ChangeVmImageDetailParam struct {
	rest string `json:"vmInstanceUuid" validate:"required"` // 必填
	rest string `json:"imageUuid" validate:"required"` // 必填
	rest string `json:"resourceUuid,omitempty"`
	rest []string `json:"tagUuids,omitempty"`
}

// ChangeVmImageParam ChangeVmImage请求参数
type ChangeVmImageParam struct {
	BaseParam
	Params ChangeVmImageDetailParam `json:"params"` // 详细参数
}

