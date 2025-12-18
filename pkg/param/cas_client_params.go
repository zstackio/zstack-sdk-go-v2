// Copyright (c) ZStack.io, Inc.

package param

// CreateCasClientDetailParam CreateCasClient详细参数
type CreateCasClientDetailParam struct {
	rest string `json:"name" validate:"required"` // 必填
	rest string `json:"description,omitempty"`
	rest string `json:"casServerLoginUrl" validate:"required"` // 必填
	rest string `json:"casServerUrlPrefix" validate:"required"` // 必填
	rest string `json:"serverName" validate:"required"` // 必填
	rest string `json:"loginType" validate:"required"` // 必填
	rest string `json:"urlTemplate" validate:"required"` // 必填
	rest []interface{} `json:"attributes,omitempty"`
	rest string `json:"resourceUuid,omitempty"`
	rest []string `json:"tagUuids,omitempty"`
}

// CreateCasClientParam CreateCasClient请求参数
type CreateCasClientParam struct {
	BaseParam
	Params CreateCasClientDetailParam `json:"params"` // 详细参数
}

