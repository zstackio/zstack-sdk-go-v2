// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// CreateHostNetworkServiceType creates HostNetworkServiceType
func (cli *ZSClient) CreateHostNetworkServiceType(params param.CreateHostNetworkServiceTypeParam) (*view.CreateHostNetworkServiceTypeEventView, error) {
	resp := view.CreateHostNetworkServiceTypeEventView{}
	if err := cli.Post("v1/hosts/service-types", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
