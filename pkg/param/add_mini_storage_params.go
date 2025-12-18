// Copyright (c) ZStack.io, Inc.

package param

// AddMiniStorageDetailParam AddMiniStorage detail param
type AddMiniStorageDetailParam struct {
	DiskIdentifier string `json:"diskIdentifier" validate:"required"`
	Url string `json:"url,omitempty"`
	Name string `json:"name" validate:"required"`
	Description string `json:"description,omitempty"`
	Type string `json:"type,omitempty"`
	ZoneUuid string `json:"zoneUuid" validate:"required"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// AddMiniStorageParam AddMiniStorage request param
type AddMiniStorageParam struct {
	BaseParam
	Params AddMiniStorageDetailParam `json:"params"`
}
