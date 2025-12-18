// Copyright (c) ZStack.io, Inc.

package param

// AddSharedBlockGroupPrimaryStorageDetailParam AddSharedBlockGroupPrimaryStorage detail param
type AddSharedBlockGroupPrimaryStorageDetailParam struct {
	DiskUuids []string `json:"diskUuids" validate:"required"`
	Url string `json:"url,omitempty"`
	Name string `json:"name" validate:"required"`
	Description string `json:"description,omitempty"`
	Type string `json:"type,omitempty"`
	ZoneUuid string `json:"zoneUuid" validate:"required"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// AddSharedBlockGroupPrimaryStorageParam AddSharedBlockGroupPrimaryStorage request param
type AddSharedBlockGroupPrimaryStorageParam struct {
	BaseParam
	Params AddSharedBlockGroupPrimaryStorageDetailParam `json:"params"`
}
