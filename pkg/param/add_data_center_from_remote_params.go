// Copyright (c) ZStack.io, Inc.

package param

// AddDataCenterFromRemoteDetailParam AddDataCenterFromRemote详细参数
type AddDataCenterFromRemoteDetailParam struct {
	rest string `json:"regionId" validate:"required"` // 必填
	rest string `json:"type" validate:"required"` // 必填
	rest bool `json:"syncZones,omitempty"`
	rest string `json:"endpoint,omitempty"`
	rest string `json:"description,omitempty"`
	rest string `json:"resourceUuid,omitempty"`
	rest []string `json:"tagUuids,omitempty"`
}

// AddDataCenterFromRemoteParam AddDataCenterFromRemote请求参数
type AddDataCenterFromRemoteParam struct {
	BaseParam
	Params AddDataCenterFromRemoteDetailParam `json:"params"` // 详细参数
}

