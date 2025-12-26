// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// UpdateHostNetworkServiceType updates HostNetworkServiceType
func (cli *ZSClient) UpdateHostNetworkServiceType(uuid string, params param.UpdateHostNetworkServiceTypeParam) (*view.UpdateHostNetworkServiceTypeEventView, error) {
	resp := view.UpdateHostNetworkServiceTypeEventView{}
	if err := cli.Put("v1/hosts/service-types/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
