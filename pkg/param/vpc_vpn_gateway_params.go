// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now // avoid unused import

// UpdateVpcVpnGatewayParamDetail UpdateVpcVpnGateway detail param
type UpdateVpcVpnGatewayParamDetail struct {
	Name string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
}

// UpdateVpcVpnGatewayParam UpdateVpcVpnGateway request param
type UpdateVpcVpnGatewayParam struct {
	BaseParam
	Params UpdateVpcVpnGatewayParamDetail `json:"updateVpcVpnGateway"`
}
