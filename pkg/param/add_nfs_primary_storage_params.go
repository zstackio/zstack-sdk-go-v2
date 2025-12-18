// Copyright (c) ZStack.io, Inc.

package param

// AddNfsPrimaryStorageDetailParam AddNfsPrimaryStorage detail param
type AddNfsPrimaryStorageDetailParam struct {
	Url string `json:"url" validate:"required"`
	Name string `json:"name" validate:"required"`
	Description string `json:"description,omitempty"`
	Type string `json:"type,omitempty"`
	ZoneUuid string `json:"zoneUuid" validate:"required"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// AddNfsPrimaryStorageParam AddNfsPrimaryStorage request param
type AddNfsPrimaryStorageParam struct {
	BaseParam
	Params AddNfsPrimaryStorageDetailParam `json:"params"`
}
