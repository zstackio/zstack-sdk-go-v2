// Copyright (c) ZStack.io, Inc.

package param

// IsOpensourceVersionDetailParam IsOpensourceVersion detail param
type IsOpensourceVersionDetailParam struct {
}

// IsOpensourceVersionParam IsOpensourceVersion request param
type IsOpensourceVersionParam struct {
	BaseParam
	Params IsOpensourceVersionDetailParam `json:"params"`
}
