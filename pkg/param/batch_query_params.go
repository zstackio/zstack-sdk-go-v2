// Copyright (c) ZStack.io, Inc.

package param

// BatchQueryDetailParam BatchQuery detail param
type BatchQueryDetailParam struct {
	Script string `json:"script,omitempty"`
}

// BatchQueryParam BatchQuery request param
type BatchQueryParam struct {
	BaseParam
	Params BatchQueryDetailParam `json:"params"`
}
