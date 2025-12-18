// Copyright (c) ZStack.io, Inc.

package param

// GetAliyunNasAccessGroupRemoteDetailParam GetAliyunNasAccessGroupRemote detail param
type GetAliyunNasAccessGroupRemoteDetailParam struct {
	DataCenterUuid string `json:"dataCenterUuid" validate:"required"`
	GroupName string `json:"groupName,omitempty"`
}

// GetAliyunNasAccessGroupRemoteParam GetAliyunNasAccessGroupRemote request param
type GetAliyunNasAccessGroupRemoteParam struct {
	BaseParam
	Params GetAliyunNasAccessGroupRemoteDetailParam `json:"params"`
}
