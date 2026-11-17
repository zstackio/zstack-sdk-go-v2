// Copyright (c) ZStack.io, Inc.

package client

import (
	"fmt"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// CreateLoadBalancerServerGroup creates LoadBalancerServerGroup
func (cli *ZSClient) CreateLoadBalancerServerGroup(loadBalancerUuid string, params param.CreateLoadBalancerServerGroupParam) (*view.LoadBalancerServerGroupInventoryView, error) {
	resp := view.LoadBalancerServerGroupInventoryView{}
	if err := cli.Post(fmt.Sprintf("v1/load-balancers/%s/servergroups", loadBalancerUuid), params, &resp); err != nil {
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

// GetCandidateVmNicsForLoadBalancerServerGroup gets CandidateVmNicsForLoadBalancerServerGroup by uuid
func (cli *ZSClient) GetCandidateVmNicsForLoadBalancerServerGroup() (*view.VmNicInventoryView, error) {
	var resp view.VmNicInventoryView
	if err := cli.GetWithRespKey("v1/load-balancers/servergroups/candidate-nics", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// RemoveBackendServerFromServerGroup removes BackendServerFromServerGroup
func (cli *ZSClient) RemoveBackendServerFromServerGroup(serverGroupUuid string, params param.RemoveBackendServerFromServerGroupParam) (*view.LoadBalancerServerGroupInventoryView, error) {
	resp := view.LoadBalancerServerGroupInventoryView{}
	if err := cli.PutWithRespKey("v1/load-balancers/servergroups", serverGroupUuid, "", map[string]interface{}{
		"removeBackendServerFromServerGroup": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ChangeLoadBalancerBackendServer changes LoadBalancerBackendServer
func (cli *ZSClient) ChangeLoadBalancerBackendServer(serverGroupUuid string, params param.ChangeLoadBalancerBackendServerParam) (*view.LoadBalancerServerGroupInventoryView, error) {
	resp := view.LoadBalancerServerGroupInventoryView{}
	if err := cli.PutWithRespKey("v1/load-balancers/servergroups", serverGroupUuid, "", map[string]interface{}{
		"changeLoadBalancerBackendServer": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AddBackendServerToServerGroup adds BackendServerToServerGroup
func (cli *ZSClient) AddBackendServerToServerGroup(serverGroupUuid string, params param.AddBackendServerToServerGroupParam) (*view.LoadBalancerServerGroupInventoryView, error) {
	resp := view.LoadBalancerServerGroupInventoryView{}
	if err := cli.Post(fmt.Sprintf("v1/load-balancers/servergroups/%s/backendservers", serverGroupUuid), params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
