// Copyright (c) ZStack.io, Inc.

package param

// CreateSNSWeComEndpointDetailParam CreateSNSWeComEndpoint详细参数
type CreateSNSWeComEndpointDetailParam struct {
	rest string `json:"url" validate:"required"` // 必填
	rest bool `json:"atAll,omitempty"`
	rest []string `json:"atPersonUserIds,omitempty"`
	rest map[string]string `json:"atPersonList,omitempty"`
	rest string `json:"name" validate:"required"` // 必填
	rest string `json:"description,omitempty"`
	rest string `json:"platformUuid,omitempty"`
	rest string `json:"resourceUuid,omitempty"`
	rest []string `json:"tagUuids,omitempty"`
}

// CreateSNSWeComEndpointParam CreateSNSWeComEndpoint请求参数
type CreateSNSWeComEndpointParam struct {
	BaseParam
	Params CreateSNSWeComEndpointDetailParam `json:"params"` // 详细参数
}

