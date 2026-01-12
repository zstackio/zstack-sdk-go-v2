// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// CreateLoadBalancerServerGroup creates LoadBalancerServerGroup
func (cli *ZSClient) CreateLoadBalancerServerGroup(params param.CreateLoadBalancerServerGroupParam) (*view.LoadBalancerServerGroupInventoryView, error) {
	var resp view.CreateLoadBalancerServerGroupEventView
	if err := cli.Post("v1/load-balancers/{loadBalancerUuid}/servergroups", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// QueryLoadBalancerServerGroup queries LoadBalancerServerGroup list
func (cli *ZSClient) QueryLoadBalancerServerGroup(params *param.QueryParam) ([]view.LoadBalancerServerGroupInventoryView, error) {
	var resp []view.LoadBalancerServerGroupInventoryView
	return resp, cli.List("v1/load-balancers/servergroups", params, &resp)
}

func (cli *ZSClient) GetLoadBalancerServerGroup(uuid string) (*view.LoadBalancerServerGroupInventoryView, error) {
	var resp view.LoadBalancerServerGroupInventoryView
	if err := cli.Get("v1/load-balancers/servergroups", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// DeleteLoadBalancerServerGroup deletes LoadBalancerServerGroup
func (cli *ZSClient) DeleteLoadBalancerServerGroup(uuid string, deleteMode param.DeleteMode) error {
	return cli.DeleteWithSpec("v1/load-balancers/servergroups", fmt.Sprintf(\"%s\", uuid), string(deleteMode))
}
// UpdateLoadBalancerServerGroup updates LoadBalancerServerGroup
func (cli *ZSClient) UpdateLoadBalancerServerGroup(uuid string, params param.UpdateLoadBalancerServerGroupParam) (*view.LoadBalancerServerGroupInventoryView, error) {
	var resp view.UpdateLoadBalancerServerGroupEventView
	err := cli.PutWithSpec("v1/load-balancers/servergroups", fmt.Sprintf(\"%s/actions\", uuid), params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
