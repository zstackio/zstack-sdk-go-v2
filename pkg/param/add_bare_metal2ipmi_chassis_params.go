// Copyright (c) ZStack.io, Inc.

package param

// AddBareMetal2IpmiChassisDetailParam AddBareMetal2IpmiChassis detail param
type AddBareMetal2IpmiChassisDetailParam struct {
	IpmiAddress string `json:"ipmiAddress" validate:"required"`
	IpmiPort int `json:"ipmiPort,omitempty"`
	IpmiUsername string `json:"ipmiUsername" validate:"required"`
	IpmiPassword string `json:"ipmiPassword" validate:"required"`
	Name string `json:"name" validate:"required"`
	Description string `json:"description,omitempty"`
	ClusterUuid string `json:"clusterUuid" validate:"required"`
	ProvisionType string `json:"provisionType,omitempty"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// AddBareMetal2IpmiChassisParam AddBareMetal2IpmiChassis request param
type AddBareMetal2IpmiChassisParam struct {
	BaseParam
	Params AddBareMetal2IpmiChassisDetailParam `json:"params"`
}
