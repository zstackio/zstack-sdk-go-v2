// Copyright (c) ZStack.io, Inc.

package param

// AddAliyunNasPrimaryStorageDetailParam AddAliyunNasPrimaryStorage detail param
type AddAliyunNasPrimaryStorageDetailParam struct {
	NasUuid string `json:"nasUuid" validate:"required"`
	AccessGroupUuid string `json:"accessGroupUuid" validate:"required"`
	VSwitchUuid string `json:"vSwitchUuid,omitempty"`
	Url string `json:"url" validate:"required"`
	Name string `json:"name" validate:"required"`
	Description string `json:"description,omitempty"`
	Type string `json:"type,omitempty"`
	ZoneUuid string `json:"zoneUuid" validate:"required"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// AddAliyunNasPrimaryStorageParam AddAliyunNasPrimaryStorage request param
type AddAliyunNasPrimaryStorageParam struct {
	BaseParam
	Params AddAliyunNasPrimaryStorageDetailParam `json:"params"`
}
