// Copyright (c) ZStack.io, Inc.

package param

// CreateVmCdRomDetailParam CreateVmCdRom detail param
type CreateVmCdRomDetailParam struct {
	Name string `json:"name" validate:"required"`
	VmInstanceUuid string `json:"vmInstanceUuid" validate:"required"`
	IsoUuid string `json:"isoUuid,omitempty"`
	Description string `json:"description,omitempty"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateVmCdRomParam CreateVmCdRom request param
type CreateVmCdRomParam struct {
	BaseParam
	Params CreateVmCdRomDetailParam `json:"params"`
}
