// Copyright (c) ZStack.io, Inc.

package param

// GetBareMetal2ProvisionNetworkIpAddressCapacityDetailParam GetBareMetal2ProvisionNetworkIpAddressCapacity detail param
type GetBareMetal2ProvisionNetworkIpAddressCapacityDetailParam struct {
	NetworkUuids []string `json:"networkUuids" validate:"required"`
}

// GetBareMetal2ProvisionNetworkIpAddressCapacityParam GetBareMetal2ProvisionNetworkIpAddressCapacity request param
type GetBareMetal2ProvisionNetworkIpAddressCapacityParam struct {
	BaseParam
	Params GetBareMetal2ProvisionNetworkIpAddressCapacityDetailParam `json:"params"`
}
