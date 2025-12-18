// Copyright (c) ZStack.io, Inc.

package param

// ExpungeBaremetalInstanceDetailParam ExpungeBaremetalInstance detail param
type ExpungeBaremetalInstanceDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
}

// ExpungeBaremetalInstanceParam ExpungeBaremetalInstance request param
type ExpungeBaremetalInstanceParam struct {
	BaseParam
	Params ExpungeBaremetalInstanceDetailParam `json:"params"`
}
