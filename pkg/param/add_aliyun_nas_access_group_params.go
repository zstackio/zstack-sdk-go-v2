// Copyright (c) ZStack.io, Inc.

package param

// AddAliyunNasAccessGroupDetailParam AddAliyunNasAccessGroup detail param
type AddAliyunNasAccessGroupDetailParam struct {
	DataCenterUuid string `json:"dataCenterUuid" validate:"required"`
	GroupName string `json:"groupName" validate:"required"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// AddAliyunNasAccessGroupParam AddAliyunNasAccessGroup request param
type AddAliyunNasAccessGroupParam struct {
	BaseParam
	Params AddAliyunNasAccessGroupDetailParam `json:"params"`
}
