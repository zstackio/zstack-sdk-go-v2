// Copyright (c) ZStack.io, Inc.

package param

// ZQLQueryDetailParam ZQLQuery detail param
type ZQLQueryDetailParam struct {
	Zql string `json:"zql,omitempty"`
}

// ZQLQueryParam ZQLQuery request param
type ZQLQueryParam struct {
	BaseParam
	Params ZQLQueryDetailParam `json:"params"`
}
