// Copyright (c) ZStack.io, Inc.

package param

// ZQLQueryDetailParam ZQLQuery详细参数
type ZQLQueryDetailParam struct {
	rest string `json:"zql,omitempty"`
}

// ZQLQueryParam ZQLQuery请求参数
type ZQLQueryParam struct {
	BaseParam
	Params ZQLQueryDetailParam `json:"params"` // 详细参数
}

