// Copyright (c) ZStack.io, Inc.

package param

// UpdateNfvInstProvisionConfigDetailParam UpdateNfvInstProvisionConfig detail param
type UpdateNfvInstProvisionConfigDetailParam struct {
	VmInstanceUuid string `json:"vmInstanceUuid" validate:"required"`
}

// UpdateNfvInstProvisionConfigParam UpdateNfvInstProvisionConfig request param
type UpdateNfvInstProvisionConfigParam struct {
	BaseParam
	Params UpdateNfvInstProvisionConfigDetailParam `json:"params"`
}
