// Copyright (c) ZStack.io, Inc.

package param

// AddExternalPrimaryStorageDetailParam AddExternalPrimaryStorage detail param
type AddExternalPrimaryStorageDetailParam struct {
	Identity string `json:"identity" validate:"required"`
	DefaultOutputProtocol string `json:"defaultOutputProtocol" validate:"required"`
	Config string `json:"config,omitempty"`
	Url string `json:"url" validate:"required"`
	Name string `json:"name" validate:"required"`
	Description string `json:"description,omitempty"`
	Type string `json:"type,omitempty"`
	ZoneUuid string `json:"zoneUuid" validate:"required"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// AddExternalPrimaryStorageParam AddExternalPrimaryStorage request param
type AddExternalPrimaryStorageParam struct {
	BaseParam
	Params AddExternalPrimaryStorageDetailParam `json:"params"`
}
