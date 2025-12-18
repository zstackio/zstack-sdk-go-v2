// Copyright (c) ZStack.io, Inc.

package param

// GetBareMetal2ProvisionNetworkIpAddressCapacityDetailParam GetBareMetal2ProvisionNetworkIpAddressCapacity详细参数
type GetBareMetal2ProvisionNetworkIpAddressCapacityDetailParam struct {
	rest []string `json:"networkUuids" validate:"required"` // 必填
}

// GetBareMetal2ProvisionNetworkIpAddressCapacityParam GetBareMetal2ProvisionNetworkIpAddressCapacity请求参数
type GetBareMetal2ProvisionNetworkIpAddressCapacityParam struct {
	BaseParam
	Params GetBareMetal2ProvisionNetworkIpAddressCapacityDetailParam `json:"params"` // 详细参数
}

