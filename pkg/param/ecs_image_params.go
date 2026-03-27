// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now() // avoid unused import

// UpdateEcsImageParamDetail UpdateEcsImage detail param
type UpdateEcsImageParamDetail struct {
	Description *string `json:"description,omitempty"`
	Name string `json:"name,omitempty"`
}

// UpdateEcsImageParam UpdateEcsImage request param
type UpdateEcsImageParam struct {
	BaseParam
	Params UpdateEcsImageParamDetail `json:"updateEcsImage"`
}
