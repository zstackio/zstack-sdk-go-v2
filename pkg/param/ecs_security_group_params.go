// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now // avoid unused import

// UpdateEcsSecurityGroupParamDetail UpdateEcsSecurityGroup detail param
type UpdateEcsSecurityGroupParamDetail struct {
	Description *string `json:"description,omitempty"`
	Name string `json:"name,omitempty"`
}

// UpdateEcsSecurityGroupParam UpdateEcsSecurityGroup request param
type UpdateEcsSecurityGroupParam struct {
	BaseParam
	Params UpdateEcsSecurityGroupParamDetail `json:"updateEcsSecurityGroup"`
}
