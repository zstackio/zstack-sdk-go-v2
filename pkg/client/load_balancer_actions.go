// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryLoadBalancer queries LoadBalancer list
func (cli *ZSClient) QueryLoadBalancer(params *param.QueryParam) ([]view.LoadBalancerInventoryView, error) {
	var resp []view.LoadBalancerInventoryView
	return resp, cli.List("v1/load-balancers", params, &resp)
}

func (cli *ZSClient) GetLoadBalancer(uuid string) (*view.LoadBalancerInventoryView, error) {
	var resp view.LoadBalancerInventoryView
	if err := cli.Get("v1/load-balancers", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
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
	err := cli.PutWithSpec("v1/load-balancers", fmt.Sprintf(\"%s/actions\", uuid), params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// DeleteLoadBalancer deletes LoadBalancer
func (cli *ZSClient) DeleteLoadBalancer(uuid string, deleteMode param.DeleteMode) error {
	return cli.DeleteWithSpec("v1/load-balancers", fmt.Sprintf(\"%s\", uuid), string(deleteMode))
}
// RefreshLoadBalancer operates on LoadBalancer
func (cli *ZSClient) RefreshLoadBalancer(uuid string, params param.RefreshLoadBalancerParam) (*view.LoadBalancerInventoryView, error) {
	var resp view.RefreshLoadBalancerEventView
	err := cli.PutWithSpec("v1/load-balancers", fmt.Sprintf(\"%s/actions\", uuid), params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
