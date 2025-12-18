// Copyright (c) ZStack.io, Inc.

package param

// CreateBaremetalChassisDetailParam CreateBaremetalChassis detail param
type CreateBaremetalChassisDetailParam struct {
	Name string `json:"name" validate:"required"`
	Description string `json:"description,omitempty"`
	ClusterUuid string `json:"clusterUuid" validate:"required"`
	IpmiAddress string `json:"ipmiAddress" validate:"required"`
	IpmiPort int `json:"ipmiPort,omitempty"`
	IpmiUsername string `json:"ipmiUsername" validate:"required"`
	IpmiPassword string `json:"ipmiPassword" validate:"required"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateBaremetalChassisParam CreateBaremetalChassis request param
type CreateBaremetalChassisParam struct {
	BaseParam
	Params CreateBaremetalChassisDetailParam `json:"params"`
}
