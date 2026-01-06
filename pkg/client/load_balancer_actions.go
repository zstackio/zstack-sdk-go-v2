// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryLoadBalancer queries LoadBalancer list
func (cli *ZSClient) QueryLoadBalancer(params *param.QueryParam) ([]view.LoadBalancerInventoryView, error) {
	var resp []view.LoadBalancerInventoryView
	return resp, cli.List("v1/load-balancers", params, &resp)
}
// CreateLoadBalancer creates LoadBalancer
func (cli *ZSClient) CreateLoadBalancer(params param.CreateLoadBalancerParam) (*view.LoadBalancerInventoryView, error) {
	var resp view.CreateLoadBalancerEventView
	if err := cli.Post("v1/load-balancers", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// UpdateLoadBalancer updates LoadBalancer
func (cli *ZSClient) UpdateLoadBalancer(uuid string, params param.UpdateLoadBalancerParam) (*view.LoadBalancerInventoryView, error) {
	var resp view.UpdateLoadBalancerEventView
	if err := cli.Put("v1/load-balancers/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// DeleteLoadBalancer deletes LoadBalancer
func (cli *ZSClient) DeleteLoadBalancer(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/load-balancers/{uuid}", uuid, string(deleteMode))
}
// RefreshLoadBalancer operates on LoadBalancer
func (cli *ZSClient) RefreshLoadBalancer(uuid string, params param.RefreshLoadBalancerParam) (*view.LoadBalancerInventoryView, error) {
	var resp view.RefreshLoadBalancerEventView
	if err := cli.Put("v1/load-balancers/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
