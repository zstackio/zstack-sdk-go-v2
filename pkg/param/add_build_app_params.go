// Copyright (c) ZStack.io, Inc.

package param

// AddBuildAppDetailParam AddBuildApp详细参数
type AddBuildAppDetailParam struct {
	rest string `json:"url" validate:"required"` // 必填
	rest string `json:"type,omitempty"`
	rest string `json:"backupStorageUuid,omitempty"`
	rest string `json:"hostname,omitempty"`
	rest string `json:"resourceUuid,omitempty"`
	rest []string `json:"tagUuids,omitempty"`
}

// AddBuildAppParam AddBuildApp请求参数
type AddBuildAppParam struct {
	BaseParam
	Params AddBuildAppDetailParam `json:"params"` // 详细参数
}

