// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// ChangeLoadBalancerBackendServer changes LoadBalancerBackendServer
func (cli *ZSClient) ChangeLoadBalancerBackendServer(uuid string, params param.ChangeLoadBalancerBackendServerParam) (*view.ChangeLoadBalancerBackendServerEventView, error) {
	resp := view.ChangeLoadBalancerBackendServerEventView{}
	if err := cli.Put("v1/load-balancers/servergroups/{serverGroupUuid}/backendserver/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
