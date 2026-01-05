// Copyright (c) ZStack.io, Inc.

package param

// CreateNfvInstGroupDetailParam CreateNfvInstGroup detail param
type CreateNfvInstGroupDetailParam struct {
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
	Params CreateNfvInstGroupDetailParam `json:"params"`
}
