// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now // avoid unused import

// GetHostNetworkInterfaceLldpParamDetail GetHostNetworkInterfaceLldp detail param
type GetHostNetworkInterfaceLldpParamDetail struct {
	InterfaceUuid string `json:"interfaceUuid" validate:"required"`
}

// GetHostNetworkInterfaceLldpParam GetHostNetworkInterfaceLldp request param
type GetHostNetworkInterfaceLldpParam struct {
	BaseParam
	Params GetHostNetworkInterfaceLldpParamDetail `json:"getHostNetworkInterfaceLldp"`
}
