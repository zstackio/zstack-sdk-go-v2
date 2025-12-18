// Copyright (c) ZStack.io, Inc.

package param

// TriggerGCJobDetailParam TriggerGCJob detail param
type TriggerGCJobDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
}

// TriggerGCJobParam TriggerGCJob request param
type TriggerGCJobParam struct {
	BaseParam
	Params TriggerGCJobDetailParam `json:"params"`
}
