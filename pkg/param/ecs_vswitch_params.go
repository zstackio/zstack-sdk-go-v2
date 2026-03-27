// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now() // avoid unused import

// UpdateEcsVSwitchParamDetail UpdateEcsVSwitch detail param
type UpdateEcsVSwitchParamDetail struct {
	Name string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
}

// UpdateEcsVSwitchParam UpdateEcsVSwitch request param
type UpdateEcsVSwitchParam struct {
	BaseParam
	Params UpdateEcsVSwitchParamDetail `json:"updateEcsVSwitch"`
}
