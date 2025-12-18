// Copyright (c) ZStack.io, Inc.

package param

// RecoverBaremetalInstanceDetailParam RecoverBaremetalInstance detail param
type RecoverBaremetalInstanceDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
}

// RecoverBaremetalInstanceParam RecoverBaremetalInstance request param
type RecoverBaremetalInstanceParam struct {
	BaseParam
	Params RecoverBaremetalInstanceDetailParam `json:"params"`
}
