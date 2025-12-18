// Copyright (c) ZStack.io, Inc.

package param

// ExpungeVmInstanceDetailParam ExpungeVmInstance detail param
type ExpungeVmInstanceDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
}

// ExpungeVmInstanceParam ExpungeVmInstance request param
type ExpungeVmInstanceParam struct {
	BaseParam
	Params ExpungeVmInstanceDetailParam `json:"params"`
}
