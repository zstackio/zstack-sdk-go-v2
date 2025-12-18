// Copyright (c) ZStack.io, Inc.

package param

// ExecuteDRSSchedulingDetailParam ExecuteDRSScheduling详细参数
type ExecuteDRSSchedulingDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
}

// ExecuteDRSSchedulingParam ExecuteDRSScheduling请求参数
type ExecuteDRSSchedulingParam struct {
	BaseParam
	Params ExecuteDRSSchedulingDetailParam `json:"params"` // 详细参数
}

