// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now // avoid unused import

// CreateNfvInstParamDetail CreateNfvInst detail param
type CreateNfvInstParamDetail struct {
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
	Params CreateNfvInstParamDetail `json:"params"`
}
// ReconnectNfvInstParamDetail ReconnectNfvInst detail param
type ReconnectNfvInstParamDetail struct {
	VmInstanceUuid string `json:"vmInstanceUuid" validate:"required"`
}

// ReconnectNfvInstParam ReconnectNfvInst request param
type ReconnectNfvInstParam struct {
	BaseParam
	Params ReconnectNfvInstParamDetail `json:"params"`
}
