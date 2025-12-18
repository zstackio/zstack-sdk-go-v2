// Copyright (c) ZStack.io, Inc.

package param

// AddLocalPrimaryStorageDetailParam AddLocalPrimaryStorage detail param
type AddLocalPrimaryStorageDetailParam struct {
	Url string `json:"url" validate:"required"`
	Name string `json:"name" validate:"required"`
	Description string `json:"description,omitempty"`
	Type string `json:"type,omitempty"`
	ZoneUuid string `json:"zoneUuid" validate:"required"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// AddLocalPrimaryStorageParam AddLocalPrimaryStorage request param
type AddLocalPrimaryStorageParam struct {
	BaseParam
	Params AddLocalPrimaryStorageDetailParam `json:"params"`
}
