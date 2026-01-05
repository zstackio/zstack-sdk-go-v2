// Copyright (c) ZStack.io, Inc.

package param

// ProvisionNfvInstGroupDetailParam ProvisionNfvInstGroup detail param
type ProvisionNfvInstGroupDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
}

// ProvisionNfvInstGroupParam ProvisionNfvInstGroup request param
type ProvisionNfvInstGroupParam struct {
	BaseParam
	Params ProvisionNfvInstGroupDetailParam `json:"params"`
}
