// Copyright (c) ZStack.io, Inc.

package param

// CreateSNSAliyunSmsEndpointDetailParam CreateSNSAliyunSmsEndpoint详细参数
type CreateSNSAliyunSmsEndpointDetailParam struct {
	rest string `json:"accessKeyUuid" validate:"required"` // 必填
	rest []string `json:"receivers,omitempty"`
	rest string `json:"name" validate:"required"` // 必填
	rest string `json:"description,omitempty"`
	rest string `json:"platformUuid,omitempty"`
	rest string `json:"resourceUuid,omitempty"`
	rest []string `json:"tagUuids,omitempty"`
}

// CreateSNSAliyunSmsEndpointParam CreateSNSAliyunSmsEndpoint请求参数
type CreateSNSAliyunSmsEndpointParam struct {
	BaseParam
	Params CreateSNSAliyunSmsEndpointDetailParam `json:"params"` // 详细参数
}

