// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now // avoid unused import

// InspectBareMetal2ChassisParamDetail InspectBareMetal2Chassis detail param
type InspectBareMetal2ChassisParamDetail struct {
}

// InspectBareMetal2ChassisParam InspectBareMetal2Chassis request param
type InspectBareMetal2ChassisParam struct {
	BaseParam
	Params InspectBareMetal2ChassisParamDetail `json:"inspectBareMetal2Chassis"`
}
// UpdateBareMetal2ChassisParamDetail UpdateBareMetal2Chassis detail param
type UpdateBareMetal2ChassisParamDetail struct {
	Name string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
}

// UpdateBareMetal2ChassisParam UpdateBareMetal2Chassis request param
type UpdateBareMetal2ChassisParam struct {
	BaseParam
	Params UpdateBareMetal2ChassisParamDetail `json:"updateBareMetal2Chassis"`
}
// DeleteBareMetal2ChassisParamDetail DeleteBareMetal2Chassis detail param
type DeleteBareMetal2ChassisParamDetail struct {
	DeleteMode *string `json:"deleteMode,omitempty"`
}

// DeleteBareMetal2ChassisParam DeleteBareMetal2Chassis request param
type DeleteBareMetal2ChassisParam struct {
	BaseParam
	Params DeleteBareMetal2ChassisParamDetail `json:"deleteBareMetal2Chassis"`
}
