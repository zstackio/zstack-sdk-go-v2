// Copyright (c) ZStack.io, Inc.

package param

// UpdatePriorityConfigDetailParam UpdatePriorityConfig detail param
type UpdatePriorityConfigDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	CpuShares int `json:"cpuShares,omitempty"`
	OomScoreAdj int `json:"oomScoreAdj,omitempty"`
}

// UpdatePriorityConfigParam UpdatePriorityConfig request param
type UpdatePriorityConfigParam struct {
	BaseParam
	Params UpdatePriorityConfigDetailParam `json:"params"`
}
