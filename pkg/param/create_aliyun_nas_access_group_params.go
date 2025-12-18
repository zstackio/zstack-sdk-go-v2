// Copyright (c) ZStack.io, Inc.

package param

// CreateAliyunNasAccessGroupDetailParam CreateAliyunNasAccessGroup detail param
type CreateAliyunNasAccessGroupDetailParam struct {
	DataCenterUuid string `json:"dataCenterUuid" validate:"required"`
	Name string `json:"name" validate:"required"`
	Description string `json:"description,omitempty"`
	NetworkType string `json:"networkType,omitempty"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateAliyunNasAccessGroupParam CreateAliyunNasAccessGroup request param
type CreateAliyunNasAccessGroupParam struct {
	BaseParam
	Params CreateAliyunNasAccessGroupDetailParam `json:"params"`
}
