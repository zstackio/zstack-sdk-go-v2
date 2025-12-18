// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// CreateHostNetworkServiceType creates HostNetworkServiceType
func (cli *ZSClient) CreateHostNetworkServiceType(params param.CreateHostNetworkServiceTypeParam) (*view.CreateHostNetworkServiceTypeEventView, error) {
	resp := view.CreateHostNetworkServiceTypeEventView{}
	if err := cli.Post("v1/hosts/service-types", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
