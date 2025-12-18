// Copyright (c) ZStack.io, Inc.

package param

// DeleteGCJobDetailParam DeleteGCJob detail param
type DeleteGCJobDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
}

// DeleteGCJobParam DeleteGCJob request param
type DeleteGCJobParam struct {
	BaseParam
	Params DeleteGCJobDetailParam `json:"params"`
}
