// Copyright (c) ZStack.io, Inc.

package param

// CreateDirectoryDetailParam CreateDirectory detail param
type CreateDirectoryDetailParam struct {
	Name string `json:"name" validate:"required"`
	ParentUuid string `json:"parentUuid,omitempty"`
	ZoneUuid string `json:"zoneUuid" validate:"required"`
	Type string `json:"type" validate:"required"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateDirectoryParam CreateDirectory request param
type CreateDirectoryParam struct {
	BaseParam
	Params CreateDirectoryDetailParam `json:"params"`
}
