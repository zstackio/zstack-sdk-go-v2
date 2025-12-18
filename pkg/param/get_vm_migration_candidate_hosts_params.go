// Copyright (c) ZStack.io, Inc.

package param

// GetVmMigrationCandidateHostsDetailParam GetVmMigrationCandidateHosts detail param
type GetVmMigrationCandidateHostsDetailParam struct {
	VmInstanceUuid string `json:"vmInstanceUuid" validate:"required"`
}

// GetVmMigrationCandidateHostsParam GetVmMigrationCandidateHosts request param
type GetVmMigrationCandidateHostsParam struct {
	BaseParam
	Params GetVmMigrationCandidateHostsDetailParam `json:"params"`
}
