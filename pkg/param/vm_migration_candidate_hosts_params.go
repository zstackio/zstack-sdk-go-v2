// Copyright (c) ZStack.io, Inc.

package param

// GetVmMigrationCandidateHostsDetailParam GetVmMigrationCandidateHosts详细参数
type GetVmMigrationCandidateHostsDetailParam struct {
	rest string `json:"vmInstanceUuid" validate:"required"` // 必填
}

// GetVmMigrationCandidateHostsParam GetVmMigrationCandidateHosts请求参数
type GetVmMigrationCandidateHostsParam struct {
	BaseParam
	Params GetVmMigrationCandidateHostsDetailParam `json:"params"` // 详细参数
}

