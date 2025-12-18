// Copyright (c) ZStack.io, Inc.

package param

// AddBlockPrimaryStorageDetailParam AddBlockPrimaryStorage detail param
type AddBlockPrimaryStorageDetailParam struct {
	VendorName string `json:"vendorName" validate:"required"`
	Metadata string `json:"metadata" validate:"required"`
	Url string `json:"url" validate:"required"`
	Name string `json:"name" validate:"required"`
	Description string `json:"description,omitempty"`
	Type string `json:"type,omitempty"`
	ZoneUuid string `json:"zoneUuid" validate:"required"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// AddBlockPrimaryStorageParam AddBlockPrimaryStorage request param
type AddBlockPrimaryStorageParam struct {
	BaseParam
	Params AddBlockPrimaryStorageDetailParam `json:"params"`
}
