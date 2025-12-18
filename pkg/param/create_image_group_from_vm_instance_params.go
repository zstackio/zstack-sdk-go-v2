// Copyright (c) ZStack.io, Inc.

package param

// CreateImageGroupFromVmInstanceDetailParam CreateImageGroupFromVmInstance detail param
type CreateImageGroupFromVmInstanceDetailParam struct {
	VmInstanceUuid string `json:"vmInstanceUuid" validate:"required"`
	Name string `json:"name" validate:"required"`
	Description string `json:"description,omitempty"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateImageGroupFromVmInstanceParam CreateImageGroupFromVmInstance request param
type CreateImageGroupFromVmInstanceParam struct {
	BaseParam
	Params CreateImageGroupFromVmInstanceDetailParam `json:"params"`
}
