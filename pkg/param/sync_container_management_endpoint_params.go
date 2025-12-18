// Copyright (c) ZStack.io, Inc.

package param

// SyncContainerManagementEndpointDetailParam SyncContainerManagementEndpoint详细参数
type SyncContainerManagementEndpointDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"zoneUuid" validate:"required"` // 必填
}

// SyncContainerManagementEndpointParam SyncContainerManagementEndpoint请求参数
type SyncContainerManagementEndpointParam struct {
	BaseParam
	Params SyncContainerManagementEndpointDetailParam `json:"params"` // 详细参数
}

