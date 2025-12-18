// Copyright (c) ZStack.io, Inc.

package param

// MigrateVmDetailParam MigrateVm detail param
type MigrateVmDetailParam struct {
	VmInstanceUuid string `json:"vmInstanceUuid" validate:"required"`
	HostUuid string `json:"hostUuid,omitempty"`
	MigrateFromDestination bool `json:"migrateFromDestination,omitempty"`
	AllowUnknown bool `json:"allowUnknown,omitempty"`
	Strategy string `json:"strategy,omitempty"`
	DownTime int `json:"downTime,omitempty"`
}

// MigrateVmParam MigrateVm request param
type MigrateVmParam struct {
	BaseParam
	Params MigrateVmDetailParam `json:"params"`
}
