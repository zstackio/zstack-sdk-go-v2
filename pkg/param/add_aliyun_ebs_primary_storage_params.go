// Copyright (c) ZStack.io, Inc.

package param

// AddAliyunEbsPrimaryStorageDetailParam AddAliyunEbsPrimaryStorage detail param
type AddAliyunEbsPrimaryStorageDetailParam struct {
	PanguPartitionUuid string `json:"panguPartitionUuid,omitempty"`
	IdentityZoneUuid string `json:"identityZoneUuid,omitempty"`
	DefaultIoType string `json:"defaultIoType,omitempty"`
	TdcConfigContent string `json:"tdcConfigContent" validate:"required"`
	Url string `json:"url" validate:"required"`
	Name string `json:"name" validate:"required"`
	Description string `json:"description,omitempty"`
	Type string `json:"type,omitempty"`
	ZoneUuid string `json:"zoneUuid" validate:"required"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// AddAliyunEbsPrimaryStorageParam AddAliyunEbsPrimaryStorage request param
type AddAliyunEbsPrimaryStorageParam struct {
	BaseParam
	Params AddAliyunEbsPrimaryStorageDetailParam `json:"params"`
}
