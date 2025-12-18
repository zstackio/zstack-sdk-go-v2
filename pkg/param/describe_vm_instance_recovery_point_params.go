// Copyright (c) ZStack.io, Inc.

package param

// DescribeVmInstanceRecoveryPointDetailParam DescribeVmInstanceRecoveryPoint详细参数
type DescribeVmInstanceRecoveryPointDetailParam struct {
	rest string `json:"vmInstanceUuid" validate:"required"` // 必填
	rest int64 `json:"groupId" validate:"required"` // 必填
	rest int `json:"limit,omitempty"`
	rest int `json:"start,omitempty"`
}

// DescribeVmInstanceRecoveryPointParam DescribeVmInstanceRecoveryPoint请求参数
type DescribeVmInstanceRecoveryPointParam struct {
	BaseParam
	Params DescribeVmInstanceRecoveryPointDetailParam `json:"params"` // 详细参数
}

