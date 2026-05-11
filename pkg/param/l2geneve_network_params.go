// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now() // avoid unused import

// CreateL2GeneveNetworkParamDetail CreateL2GeneveNetwork detail param
type CreateL2GeneveNetworkParamDetail struct {
	GeneveId int `json:"geneveId" validate:"required"`
	Name string `json:"name" validate:"required"`
	Description *string `json:"description,omitempty"`
	ZoneUuid string `json:"zoneUuid" validate:"required"`
	PhysicalInterface *string `json:"physicalInterface,omitempty"`
	Type *string `json:"type,omitempty"`
	VSwitchType *string `json:"vSwitchType,omitempty"`
	Isolated *bool `json:"isolated,omitempty"`
	Pvlan *string `json:"pvlan,omitempty"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateL2GeneveNetworkParam CreateL2GeneveNetwork request param
type CreateL2GeneveNetworkParam struct {
	BaseParam
	Params CreateL2GeneveNetworkParamDetail `json:"params"`
}
