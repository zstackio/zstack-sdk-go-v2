// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now // avoid unused import

// AddAliyunNasAccessGroupParamDetail AddAliyunNasAccessGroup detail param
type AddAliyunNasAccessGroupParamDetail struct {
	DataCenterUuid string `json:"dataCenterUuid" validate:"required"`
	GroupName string `json:"groupName" validate:"required"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// AddAliyunNasAccessGroupParam AddAliyunNasAccessGroup request param
type AddAliyunNasAccessGroupParam struct {
	BaseParam
	AddAliyunNasAccessGroup AddAliyunNasAccessGroupParamDetail `json:"addAliyunNasAccessGroup"`
}
// UpdateAliyunNasAccessGroupParamDetail UpdateAliyunNasAccessGroup detail param
type UpdateAliyunNasAccessGroupParamDetail struct {
	Uuid string `json:"uuid" validate:"required"`
	Description string `json:"description,omitempty"`
}

// UpdateAliyunNasAccessGroupParam UpdateAliyunNasAccessGroup request param
type UpdateAliyunNasAccessGroupParam struct {
	BaseParam
	UpdateAliyunNasAccessGroup UpdateAliyunNasAccessGroupParamDetail `json:"updateAliyunNasAccessGroup"`
}
// CreateAliyunNasAccessGroupParamDetail CreateAliyunNasAccessGroup detail param
type CreateAliyunNasAccessGroupParamDetail struct {
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
	CreateAliyunNasAccessGroup CreateAliyunNasAccessGroupParamDetail `json:"createAliyunNasAccessGroup"`
}
// DeleteAliyunNasAccessGroupParamDetail DeleteAliyunNasAccessGroup detail param
type DeleteAliyunNasAccessGroupParamDetail struct {
	Uuid string `json:"uuid" validate:"required"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// DeleteAliyunNasAccessGroupParam DeleteAliyunNasAccessGroup request param
type DeleteAliyunNasAccessGroupParam struct {
	BaseParam
	DeleteAliyunNasAccessGroup DeleteAliyunNasAccessGroupParamDetail `json:"deleteAliyunNasAccessGroup"`
}
