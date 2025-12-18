// Copyright (c) ZStack.io, Inc.

package param

// AddContainerManagementEndpointDetailParam AddContainerManagementEndpoint详细参数
type AddContainerManagementEndpointDetailParam struct {
	rest string `json:"name" validate:"required"` // 必填
	rest string `json:"description,omitempty"`
	rest string `json:"managementIp" validate:"required"` // 必填
	rest string `json:"vendor" validate:"required"` // 必填
	rest int `json:"managementPort" validate:"required"` // 必填
	rest string `json:"containerAccessKeyId" validate:"required"` // 必填
	rest string `json:"containerAccessKeySecret" validate:"required"` // 必填
	rest string `json:"resourceUuid,omitempty"`
	rest []string `json:"tagUuids,omitempty"`
}

// AddContainerManagementEndpointParam AddContainerManagementEndpoint请求参数
type AddContainerManagementEndpointParam struct {
	BaseParam
	Params AddContainerManagementEndpointDetailParam `json:"params"` // 详细参数
}

