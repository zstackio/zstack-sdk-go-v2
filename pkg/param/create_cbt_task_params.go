// Copyright (c) ZStack.io, Inc.

package param

// CreateCbtTaskDetailParam CreateCbtTask detail param
type CreateCbtTaskDetailParam struct {
	Name string `json:"name" validate:"required"`
	Description string `json:"description,omitempty"`
	VmInstanceUuid string `json:"vmInstanceUuid" validate:"required"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateCbtTaskParam CreateCbtTask request param
type CreateCbtTaskParam struct {
	BaseParam
	Params CreateCbtTaskDetailParam `json:"params"`
}
