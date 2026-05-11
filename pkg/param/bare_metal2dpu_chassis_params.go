// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now() // avoid unused import

// AddBareMetal2DpuChassisParamDetail AddBareMetal2DpuChassis detail param
type AddBareMetal2DpuChassisParamDetail struct {
	Url string `json:"url" validate:"required"`
	VendorType string `json:"vendorType" validate:"required"`
	Config *string `json:"config,omitempty"`
	Name string `json:"name" validate:"required"`
	Description *string `json:"description,omitempty"`
	ClusterUuid string `json:"clusterUuid" validate:"required"`
	ProvisionType *string `json:"provisionType,omitempty"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// AddBareMetal2DpuChassisParam AddBareMetal2DpuChassis request param
type AddBareMetal2DpuChassisParam struct {
	BaseParam
	Params AddBareMetal2DpuChassisParamDetail `json:"params"`
}
