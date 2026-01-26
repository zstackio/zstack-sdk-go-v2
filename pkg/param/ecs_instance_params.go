// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now // avoid unused import

// StartEcsInstanceParamDetail StartEcsInstance detail param
type StartEcsInstanceParamDetail struct {
}

// StartEcsInstanceParam StartEcsInstance request param
type StartEcsInstanceParam struct {
	BaseParam
	Params StartEcsInstanceParamDetail `json:"startEcsInstance"`
}
// DeleteEcsInstanceParamDetail DeleteEcsInstance detail param
type DeleteEcsInstanceParamDetail struct {
	DeleteMode *string `json:"deleteMode,omitempty"`
}

// DeleteEcsInstanceParam DeleteEcsInstance request param
type DeleteEcsInstanceParam struct {
	BaseParam
	Params DeleteEcsInstanceParamDetail `json:"deleteEcsInstance"`
}
// StopEcsInstanceParamDetail StopEcsInstance detail param
type StopEcsInstanceParamDetail struct {
}

// StopEcsInstanceParam StopEcsInstance request param
type StopEcsInstanceParam struct {
	BaseParam
	Params StopEcsInstanceParamDetail `json:"stopEcsInstance"`
}
// RebootEcsInstanceParamDetail RebootEcsInstance detail param
type RebootEcsInstanceParamDetail struct {
}

// RebootEcsInstanceParam RebootEcsInstance request param
type RebootEcsInstanceParam struct {
	BaseParam
	Params RebootEcsInstanceParamDetail `json:"rebootEcsInstance"`
}
// UpdateEcsInstanceParamDetail UpdateEcsInstance detail param
type UpdateEcsInstanceParamDetail struct {
	Name string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	Password *string `json:"password,omitempty"`
}

// UpdateEcsInstanceParam UpdateEcsInstance request param
type UpdateEcsInstanceParam struct {
	BaseParam
	Params UpdateEcsInstanceParamDetail `json:"params"`
}
