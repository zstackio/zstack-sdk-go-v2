// Copyright (c) ZStack.io, Inc.

package param

// GetVmTaskDetailParam GetVmTask detail param
type GetVmTaskDetailParam struct {
	VmInstanceUuids []string `json:"vmInstanceUuids" validate:"required"`
	SyncSignatures []string `json:"syncSignatures,omitempty"`
}

// GetVmTaskParam GetVmTask request param
type GetVmTaskParam struct {
	BaseParam
	Params GetVmTaskDetailParam `json:"params"`
}
