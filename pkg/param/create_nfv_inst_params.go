// Copyright (c) ZStack.io, Inc.

package param

// CreateNfvInstDetailParam CreateNfvInst detail param
type CreateNfvInstDetailParam struct {
	Name string `json:"name" validate:"required"`
	NfvInstGroupUuid string `json:"nfvInstGroupUuid" validate:"required"`
	Description string `json:"description,omitempty"`
	ZoneUuid string `json:"zoneUuid,omitempty"`
	ClusterUuid string `json:"clusterUuid,omitempty"`
	HostUuid string `json:"hostUuid,omitempty"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateNfvInstParam CreateNfvInst request param
type CreateNfvInstParam struct {
	BaseParam
	Params CreateNfvInstDetailParam `json:"params"`
}
