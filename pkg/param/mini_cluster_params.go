// Copyright (c) ZStack.io, Inc.

package param

// CreateMiniClusterDetailParam CreateMiniCluster详细参数
type CreateMiniClusterDetailParam struct {
	rest string `json:"zoneUuid" validate:"required"` // 必填
	rest string `json:"name" validate:"required"` // 必填
	rest []string `json:"hostManagementIps" validate:"required"` // 必填
	rest string `json:"username,omitempty"`
	rest string `json:"password" validate:"required"` // 必填
	rest int `json:"sshPort,omitempty"`
	rest string `json:"description,omitempty"`
	rest string `json:"hypervisorType" validate:"required"` // 必填
	rest string `json:"resourceUuid,omitempty"`
	rest []string `json:"tagUuids,omitempty"`
}

// CreateMiniClusterParam CreateMiniCluster请求参数
type CreateMiniClusterParam struct {
	BaseParam
	Params CreateMiniClusterDetailParam `json:"params"` // 详细参数
}

