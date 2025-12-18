// Copyright (c) ZStack.io, Inc.

package param

// CreateImageReplicationGroupDetailParam CreateImageReplicationGroup detail param
type CreateImageReplicationGroupDetailParam struct {
	Name string `json:"name" validate:"required"`
	Description string `json:"description,omitempty"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateImageReplicationGroupParam CreateImageReplicationGroup request param
type CreateImageReplicationGroupParam struct {
	BaseParam
	Params CreateImageReplicationGroupDetailParam `json:"params"`
}
