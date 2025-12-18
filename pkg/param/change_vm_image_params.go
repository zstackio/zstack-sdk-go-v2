// Copyright (c) ZStack.io, Inc.

package param

// ChangeVmImageDetailParam ChangeVmImage detail param
type ChangeVmImageDetailParam struct {
	VmInstanceUuid string `json:"vmInstanceUuid" validate:"required"`
	ImageUuid string `json:"imageUuid" validate:"required"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// ChangeVmImageParam ChangeVmImage request param
type ChangeVmImageParam struct {
	BaseParam
	Params ChangeVmImageDetailParam `json:"params"`
}
