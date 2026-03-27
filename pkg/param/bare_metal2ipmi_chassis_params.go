// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now() // avoid unused import

// AddBareMetal2IpmiChassisParamDetail AddBareMetal2IpmiChassis detail param
type AddBareMetal2IpmiChassisParamDetail struct {
	IpmiAddress string `json:"ipmiAddress" validate:"required"`
	IpmiPort *int `json:"ipmiPort,omitempty"`
	IpmiUsername string `json:"ipmiUsername" validate:"required"`
	IpmiPassword string `json:"ipmiPassword" validate:"required"`
	Name string `json:"name" validate:"required"`
	Description *string `json:"description,omitempty"`
	ClusterUuid string `json:"clusterUuid" validate:"required"`
	ProvisionType *string `json:"provisionType,omitempty"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// AddBareMetal2IpmiChassisParam AddBareMetal2IpmiChassis request param
type AddBareMetal2IpmiChassisParam struct {
	BaseParam
	Params AddBareMetal2IpmiChassisParamDetail `json:"params"`
}
// UpdateBareMetal2IpmiChassisParamDetail UpdateBareMetal2IpmiChassis detail param
type UpdateBareMetal2IpmiChassisParamDetail struct {
	IpmiAddress *string `json:"ipmiAddress,omitempty"`
	IpmiPort *int `json:"ipmiPort,omitempty"`
	IpmiUsername *string `json:"ipmiUsername,omitempty"`
	IpmiPassword *string `json:"ipmiPassword,omitempty"`
	Name string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
}

// UpdateBareMetal2IpmiChassisParam UpdateBareMetal2IpmiChassis request param
type UpdateBareMetal2IpmiChassisParam struct {
	BaseParam
	Params UpdateBareMetal2IpmiChassisParamDetail `json:"updateBareMetal2IpmiChassis"`
}
