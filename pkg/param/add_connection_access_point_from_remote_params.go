// Copyright (c) ZStack.io, Inc.

package param

// AddConnectionAccessPointFromRemoteDetailParam AddConnectionAccessPointFromRemote详细参数
type AddConnectionAccessPointFromRemoteDetailParam struct {
	rest string `json:"dataCenterUuid" validate:"required"` // 必填
	rest string `json:"accessPointId" validate:"required"` // 必填
	rest string `json:"resourceUuid,omitempty"`
	rest []string `json:"tagUuids,omitempty"`
}

// AddConnectionAccessPointFromRemoteParam AddConnectionAccessPointFromRemote请求参数
type AddConnectionAccessPointFromRemoteParam struct {
	BaseParam
	Params AddConnectionAccessPointFromRemoteDetailParam `json:"params"` // 详细参数
}

