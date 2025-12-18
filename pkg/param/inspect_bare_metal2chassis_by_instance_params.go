// Copyright (c) ZStack.io, Inc.

package param

// InspectBareMetal2ChassisByInstanceDetailParam InspectBareMetal2ChassisByInstance detail param
type InspectBareMetal2ChassisByInstanceDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
}

// InspectBareMetal2ChassisByInstanceParam InspectBareMetal2ChassisByInstance request param
type InspectBareMetal2ChassisByInstanceParam struct {
	BaseParam
	Params InspectBareMetal2ChassisByInstanceDetailParam `json:"params"`
}
