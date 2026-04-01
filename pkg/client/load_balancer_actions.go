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

// PageLoadBalancer Pagination
func (cli *ZSClient) PageLoadBalancer(params *param.QueryParam) ([]view.LoadBalancerInventoryView, int, error) {
	var loadBalancers []view.LoadBalancerInventoryView
	total, err := cli.Page("v1/load-balancers", params, &loadBalancers)
	return loadBalancers, total, err
}
// CreateLoadBalancer creates LoadBalancer
func (cli *ZSClient) CreateLoadBalancer(params param.CreateLoadBalancerParam) (*view.LoadBalancerInventoryView, error) {
	resp := view.LoadBalancerInventoryView{}
	if err := cli.Post("v1/load-balancers", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// UpdateLoadBalancer updates LoadBalancer
func (cli *ZSClient) UpdateLoadBalancer(uuid string, params param.UpdateLoadBalancerParam) (*view.LoadBalancerInventoryView, error) {
	resp := view.LoadBalancerInventoryView{}
	if err := cli.PutWithRespKey("v1/load-balancers", uuid, "", map[string]interface{}{
		"updateLoadBalancer": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// DeleteLoadBalancer deletes LoadBalancer
func (cli *ZSClient) DeleteLoadBalancer(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/load-balancers", uuid, string(deleteMode))
}
// RefreshLoadBalancer operates on LoadBalancer
func (cli *ZSClient) RefreshLoadBalancer(uuid string, params param.RefreshLoadBalancerParam) (*view.LoadBalancerInventoryView, error) {
	resp := view.LoadBalancerInventoryView{}
	if err := cli.PutWithRespKey("v1/load-balancers", uuid, "", map[string]interface{}{
		"refreshLoadBalancer": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
