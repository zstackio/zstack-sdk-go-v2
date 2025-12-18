// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// UpdateHostNetworkServiceType updates HostNetworkServiceType
func (cli *ZSClient) UpdateHostNetworkServiceType(uuid string, params param.UpdateHostNetworkServiceTypeParam) (*view.UpdateHostNetworkServiceTypeEventView, error) {
	resp := view.UpdateHostNetworkServiceTypeEventView{}
	if err := cli.Put("v1/hosts/service-types/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
