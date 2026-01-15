// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now // avoid unused import

// UpdateHostNetworkInterfaceParamDetail UpdateHostNetworkInterface detail param
type UpdateHostNetworkInterfaceParamDetail struct {
	InterfaceUuid string `json:"interfaceUuid" validate:"required"`
	Description string `json:"description" validate:"required"`
}

// UpdateHostNetworkInterfaceParam UpdateHostNetworkInterface request param
type UpdateHostNetworkInterfaceParam struct {
	BaseParam
	UpdateHostNetworkInterface UpdateHostNetworkInterfaceParamDetail `json:"updateHostNetworkInterface"`
}
// LocateHostNetworkInterfaceParamDetail LocateHostNetworkInterface detail param
type LocateHostNetworkInterfaceParamDetail struct {
	HostUuid string `json:"hostUuid" validate:"required"`
	NetworkInterfaceName string `json:"networkInterfaceName" validate:"required"`
	Interval int64 `json:"interval,omitempty"`
}

// LocateHostNetworkInterfaceParam LocateHostNetworkInterface request param
type LocateHostNetworkInterfaceParam struct {
	BaseParam
	LocateHostNetworkInterface LocateHostNetworkInterfaceParamDetail `json:"locateHostNetworkInterface"`
}
