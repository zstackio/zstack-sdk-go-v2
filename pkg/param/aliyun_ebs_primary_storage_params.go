// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now // avoid unused import

// UpdateAliyunEbsPrimaryStorageParamDetail UpdateAliyunEbsPrimaryStorage detail param
type UpdateAliyunEbsPrimaryStorageParamDetail struct {
	PanguAppName string `json:"panguAppName,omitempty"`
	PanguPartitionName string `json:"panguPartitionName,omitempty"`
	Uuid string `json:"uuid" validate:"required"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Url string `json:"url,omitempty"`
}

// UpdateAliyunEbsPrimaryStorageParam UpdateAliyunEbsPrimaryStorage request param
type UpdateAliyunEbsPrimaryStorageParam struct {
	BaseParam
	Params UpdateAliyunEbsPrimaryStorageParamDetail `json:"params"`
}
// AddAliyunEbsPrimaryStorageParamDetail AddAliyunEbsPrimaryStorage detail param
type AddAliyunEbsPrimaryStorageParamDetail struct {
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
	Params AddAliyunEbsPrimaryStorageParamDetail `json:"params"`
}
