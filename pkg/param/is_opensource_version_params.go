// Copyright (c) ZStack.io, Inc.

package param

// IsOpensourceVersionDetailParam IsOpensourceVersion详细参数
type IsOpensourceVersionDetailParam struct {
}

// IsOpensourceVersionParam IsOpensourceVersion请求参数
type IsOpensourceVersionParam struct {
	BaseParam
	Params IsOpensourceVersionDetailParam `json:"params"` // 详细参数
}

