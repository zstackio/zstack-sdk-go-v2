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
	resp := view.LoadBalancerServerGroupInventoryView{}
	if err := cli.Post("v1/load-balancers/{loadBalancerUuid}/servergroups", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
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

// PageLoadBalancerServerGroup Pagination
func (cli *ZSClient) PageLoadBalancerServerGroup(params *param.QueryParam) ([]view.LoadBalancerServerGroupInventoryView, int, error) {
	var loadBalancerServerGroups []view.LoadBalancerServerGroupInventoryView
	total, err := cli.Page("v1/load-balancers/servergroups", params, &loadBalancerServerGroups)
	return loadBalancerServerGroups, total, err
}
// DeleteLoadBalancerServerGroup deletes LoadBalancerServerGroup
func (cli *ZSClient) DeleteLoadBalancerServerGroup(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/load-balancers/servergroups", uuid, string(deleteMode))
}
// UpdateLoadBalancerServerGroup updates LoadBalancerServerGroup
func (cli *ZSClient) UpdateLoadBalancerServerGroup(uuid string, params param.UpdateLoadBalancerServerGroupParam) (*view.LoadBalancerServerGroupInventoryView, error) {
	resp := view.LoadBalancerServerGroupInventoryView{}
	if err := cli.PutWithRespKey("v1/load-balancers/servergroups", uuid, "", map[string]interface{}{
		"updateLoadBalancerServerGroup": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
