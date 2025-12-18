// Copyright (c) ZStack.io, Inc.

package param

// ExecuteDRSSchedulingDetailParam ExecuteDRSScheduling detail param
type ExecuteDRSSchedulingDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
}

// ExecuteDRSSchedulingParam ExecuteDRSScheduling request param
type ExecuteDRSSchedulingParam struct {
	BaseParam
	Params ExecuteDRSSchedulingDetailParam `json:"params"`
}
