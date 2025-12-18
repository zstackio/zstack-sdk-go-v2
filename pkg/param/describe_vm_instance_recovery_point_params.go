// Copyright (c) ZStack.io, Inc.

package param

// DescribeVmInstanceRecoveryPointDetailParam DescribeVmInstanceRecoveryPoint detail param
type DescribeVmInstanceRecoveryPointDetailParam struct {
	VmInstanceUuid string `json:"vmInstanceUuid" validate:"required"`
	GroupId int64 `json:"groupId" validate:"required"`
	Limit int `json:"limit,omitempty"`
	Start int `json:"start,omitempty"`
}

// DescribeVmInstanceRecoveryPointParam DescribeVmInstanceRecoveryPoint request param
type DescribeVmInstanceRecoveryPointParam struct {
	BaseParam
	Params DescribeVmInstanceRecoveryPointDetailParam `json:"params"`
}
