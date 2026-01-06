// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
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
// DeleteLoadBalancerServerGroup deletes LoadBalancerServerGroup
func (cli *ZSClient) DeleteLoadBalancerServerGroup(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/load-balancers/servergroups/{uuid}", uuid, string(deleteMode))
}
// UpdateLoadBalancerServerGroup updates LoadBalancerServerGroup
func (cli *ZSClient) UpdateLoadBalancerServerGroup(uuid string, params param.UpdateLoadBalancerServerGroupParam) (*view.LoadBalancerServerGroupInventoryView, error) {
	var resp view.UpdateLoadBalancerServerGroupEventView
	if err := cli.Put("v1/load-balancers/servergroups/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
