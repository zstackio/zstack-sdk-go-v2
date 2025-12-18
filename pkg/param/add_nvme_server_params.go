// Copyright (c) ZStack.io, Inc.

package param

// AddNvmeServerDetailParam AddNvmeServer详细参数
type AddNvmeServerDetailParam struct {
	rest string `json:"name,omitempty"`
	rest string `json:"ip" validate:"required"` // 必填
	rest int `json:"port,omitempty"`
	rest string `json:"transport" validate:"required"` // 必填
	rest string `json:"resourceUuid,omitempty"`
	rest []string `json:"tagUuids,omitempty"`
}

// AddNvmeServerParam AddNvmeServer请求参数
type AddNvmeServerParam struct {
	BaseParam
	Params AddNvmeServerDetailParam `json:"params"` // 详细参数
}

