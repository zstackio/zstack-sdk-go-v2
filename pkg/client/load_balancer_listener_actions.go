// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// UpdateLoadBalancerListener updates LoadBalancerListener
func (cli *ZSClient) UpdateLoadBalancerListener(uuid string, params param.UpdateLoadBalancerListenerParam) (*view.LoadBalancerListenerInventoryView, error) {
	resp := view.LoadBalancerListenerInventoryView{}
	if err := cli.PutWithRespKey("v1/load-balancers/listeners", uuid, "", map[string]interface{}{
		"updateLoadBalancerListener": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// CreateLoadBalancerListener creates LoadBalancerListener
func (cli *ZSClient) CreateLoadBalancerListener(params param.CreateLoadBalancerListenerParam) (*view.LoadBalancerListenerInventoryView, error) {
	resp := view.LoadBalancerListenerInventoryView{}
	if err := cli.Post("v1/load-balancers/{loadBalancerUuid}/listeners", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// QueryLoadBalancerListener queries LoadBalancerListener list
func (cli *ZSClient) QueryLoadBalancerListener(params *param.QueryParam) ([]view.LoadBalancerListenerInventoryView, error) {
	var resp []view.LoadBalancerListenerInventoryView
	return resp, cli.List("v1/load-balancers/listeners", params, &resp)
}

func (cli *ZSClient) GetLoadBalancerListener(uuid string) (*view.LoadBalancerListenerInventoryView, error) {
	var resp view.LoadBalancerListenerInventoryView
	if err := cli.Get("v1/load-balancers/listeners", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageLoadBalancerListener Pagination
func (cli *ZSClient) PageLoadBalancerListener(params *param.QueryParam) ([]view.LoadBalancerListenerInventoryView, int, error) {
	var loadBalancerListeners []view.LoadBalancerListenerInventoryView
	total, err := cli.Page("v1/load-balancers/listeners", params, &loadBalancerListeners)
	return loadBalancerListeners, total, err
}
// ChangeLoadBalancerListener changes LoadBalancerListener
func (cli *ZSClient) ChangeLoadBalancerListener(uuid string, params param.ChangeLoadBalancerListenerParam) (*view.LoadBalancerListenerInventoryView, error) {
	resp := view.LoadBalancerListenerInventoryView{}
	if err := cli.PutWithRespKey("v1/load-balancers/listeners", uuid, "", map[string]interface{}{
		"changeLoadBalancerListener": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// DeleteLoadBalancerListener deletes LoadBalancerListener
func (cli *ZSClient) DeleteLoadBalancerListener(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/load-balancers/listeners", uuid, string(deleteMode))
}
