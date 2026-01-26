// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now // avoid unused import

// UpdateVpcUserVpnGatewayParamDetail UpdateVpcUserVpnGateway detail param
type UpdateVpcUserVpnGatewayParamDetail struct {
	Name string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
}

// UpdateVpcUserVpnGatewayParam UpdateVpcUserVpnGateway request param
type UpdateVpcUserVpnGatewayParam struct {
	BaseParam
	Params UpdateVpcUserVpnGatewayParamDetail `json:"updateVpcUserVpnGateway"`
}
