// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now // avoid unused import

// CreateNfvInstGroupParamDetail CreateNfvInstGroup detail param
type CreateNfvInstGroupParamDetail struct {
	Name string `json:"name" validate:"required"`
	Description string `json:"description,omitempty"`
	InstType string `json:"instType" validate:"required"`
	FuncType string `json:"funcType" validate:"required"`
	MonitorIps []string `json:"monitorIps,omitempty"`
	NfvInstOfferingUuid string `json:"nfvInstOfferingUuid" validate:"required"`
	FrontEndL3NetworkUuid string `json:"frontEndL3NetworkUuid,omitempty"`
	BackendL3NetworkUuids []string `json:"backendL3NetworkUuids,omitempty"`
	NetOsDistro string `json:"netOsDistro,omitempty"`
	BaseOsDistro string `json:"baseOsDistro,omitempty"`
	VipUuid string `json:"vipUuid,omitempty"`
	Ipv6VipUuid string `json:"ipv6VipUuid,omitempty"`
	PrimaryStorageUuid string `json:"primaryStorageUuid,omitempty"`
	PrimaryStoragePoolUuid string `json:"primaryStoragePoolUuid,omitempty"`
	ClusterUuid string `json:"clusterUuid,omitempty"`
	ZoneUuid string `json:"zoneUuid,omitempty"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateNfvInstGroupParam CreateNfvInstGroup request param
type CreateNfvInstGroupParam struct {
	BaseParam
	Params CreateNfvInstGroupParamDetail `json:"createNfvInstGroup"`
}
// SyncNfvInstGroupParamDetail SyncNfvInstGroup detail param
type SyncNfvInstGroupParamDetail struct {
	Uuid string `json:"uuid" validate:"required"`
}

// SyncNfvInstGroupParam SyncNfvInstGroup request param
type SyncNfvInstGroupParam struct {
	BaseParam
	Params SyncNfvInstGroupParamDetail `json:"syncNfvInstGroup"`
}
// UpdateNfvInstGroupParamDetail UpdateNfvInstGroup detail param
type UpdateNfvInstGroupParamDetail struct {
	Uuid string `json:"uuid" validate:"required"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

// UpdateNfvInstGroupParam UpdateNfvInstGroup request param
type UpdateNfvInstGroupParam struct {
	BaseParam
	Params UpdateNfvInstGroupParamDetail `json:"updateNfvInstGroup"`
}
// DeleteNfvInstGroupParamDetail DeleteNfvInstGroup detail param
type DeleteNfvInstGroupParamDetail struct {
	Uuid string `json:"uuid" validate:"required"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// DeleteNfvInstGroupParam DeleteNfvInstGroup request param
type DeleteNfvInstGroupParam struct {
	BaseParam
	Params DeleteNfvInstGroupParamDetail `json:"deleteNfvInstGroup"`
}
