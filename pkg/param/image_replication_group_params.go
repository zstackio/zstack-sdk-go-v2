// Copyright (c) ZStack.io, Inc.

package param

// CreateImageReplicationGroupDetailParam CreateImageReplicationGroup详细参数
type CreateImageReplicationGroupDetailParam struct {
	rest string `json:"name" validate:"required"` // 必填
	rest string `json:"description,omitempty"`
	rest string `json:"resourceUuid,omitempty"`
	rest []string `json:"tagUuids,omitempty"`
}

// CreateImageReplicationGroupParam CreateImageReplicationGroup请求参数
type CreateImageReplicationGroupParam struct {
	BaseParam
	Params CreateImageReplicationGroupDetailParam `json:"params"` // 详细参数
}

