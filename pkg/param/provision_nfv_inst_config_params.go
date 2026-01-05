// Copyright (c) ZStack.io, Inc.

package param

// ProvisionNfvInstConfigDetailParam ProvisionNfvInstConfig detail param
type ProvisionNfvInstConfigDetailParam struct {
	VmInstanceUuid string `json:"vmInstanceUuid" validate:"required"`
}

// ProvisionNfvInstConfigParam ProvisionNfvInstConfig request param
type ProvisionNfvInstConfigParam struct {
	BaseParam
	Params ProvisionNfvInstConfigDetailParam `json:"params"`
}
