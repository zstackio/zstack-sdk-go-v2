// Copyright (c) ZStack.io, Inc.

package param

// LocateHostNetworkInterfaceDetailParam LocateHostNetworkInterface detail param
type LocateHostNetworkInterfaceDetailParam struct {
	HostUuid string `json:"hostUuid" validate:"required"`
	NetworkInterfaceName string `json:"networkInterfaceName" validate:"required"`
	Interval int64 `json:"interval,omitempty"`
}

// LocateHostNetworkInterfaceParam LocateHostNetworkInterface request param
type LocateHostNetworkInterfaceParam struct {
	BaseParam
	Params LocateHostNetworkInterfaceDetailParam `json:"params"`
}
