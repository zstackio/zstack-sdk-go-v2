// Copyright (c) ZStack.io, Inc.

package param

// AddSharedMountPointPrimaryStorageDetailParam AddSharedMountPointPrimaryStorage detail param
type AddSharedMountPointPrimaryStorageDetailParam struct {
	Url string `json:"url" validate:"required"`
	Name string `json:"name" validate:"required"`
	Description string `json:"description,omitempty"`
	Type string `json:"type,omitempty"`
	ZoneUuid string `json:"zoneUuid" validate:"required"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// AddSharedMountPointPrimaryStorageParam AddSharedMountPointPrimaryStorage request param
type AddSharedMountPointPrimaryStorageParam struct {
	BaseParam
	Params AddSharedMountPointPrimaryStorageDetailParam `json:"params"`
}
