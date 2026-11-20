// Copyright (c) ZStack.io, Inc.

package client

import (
	"fmt"
	"net/url"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryVirtualRouterVm queries VirtualRouterVm list
func (cli *ZSClient) QueryVirtualRouterVm(params *param.QueryParam) ([]view.ApplianceVmInventoryView, error) {
	var resp []view.ApplianceVmInventoryView
	return resp, cli.List("v1/vm-instances/appliances/virtual-routers", params, &resp)
}

func (cli *ZSClient) GetVirtualRouterVm(uuid string) (*view.ApplianceVmInventoryView, error) {
	var resp view.ApplianceVmInventoryView
	if err := cli.Get("v1/vm-instances/appliances/virtual-routers", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageVirtualRouterVm Pagination
func (cli *ZSClient) PageVirtualRouterVm(params *param.QueryParam) ([]view.ApplianceVmInventoryView, int, error) {
	var virtualRouterVms []view.ApplianceVmInventoryView
	total, err := cli.Page("v1/vm-instances/appliances/virtual-routers", params, &virtualRouterVms)
	return virtualRouterVms, total, err
}

// RemoveDnsFromVpcRouter requires additional arguments.
//
// Deprecated: use RemoveDnsFromVpcRouterWithParams to supply the required fields.
func (cli *ZSClient) RemoveDnsFromVpcRouter(uuid string, deleteMode param.DeleteMode) error {
	return fmt.Errorf("RemoveDnsFromVpcRouter requires additional arguments; use RemoveDnsFromVpcRouterWithParams")
}

// RemoveDnsFromVpcRouterWithParams sends the DELETE request with its operation parameters.
func (cli *ZSClient) RemoveDnsFromVpcRouterWithParams(uuid string, deleteMode param.DeleteMode, params param.RemoveDnsFromVpcRouterParamDetail) error {
	query := url.Values{
		"dns": {params.Dns},
	}
	return cli.deleteActionWithQuery("v1/vpc/virtual-routers", uuid, "dns", deleteMode, query, "dns")
}

// ReconnectVirtualRouter operates on VirtualRouter
func (cli *ZSClient) ReconnectVirtualRouter(vmInstanceUuid string, params param.ReconnectVirtualRouterParam) (*view.ApplianceVmInventoryView, error) {
	resp := view.ApplianceVmInventoryView{}
	if err := cli.PutWithRespKey("v1/vm-instances/appliances/virtual-routers", vmInstanceUuid, "", map[string]interface{}{
		"reconnectVirtualRouter": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetVpcVRouterDistributedRoutingConnections gets VpcVRouterDistributedRoutingConnections by uuid
func (cli *ZSClient) GetVpcVRouterDistributedRoutingConnections(uuid string) (*view.StringView, error) {
	var resp view.StringView
	if err := cli.GetWithRespKey("v1/vpc/virtual-routers", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetVpcVRouterNetworkServiceState gets VpcVRouterNetworkServiceState by uuid
func (cli *ZSClient) GetVpcVRouterNetworkServiceState(uuid string) (*view.GetVpcVRouterNetworkServiceStateView, error) {
	var resp view.GetVpcVRouterNetworkServiceStateView
	if err := cli.GetWithRespKey("v1/vpc/virtual-routers", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SetVpcVRouterNetworkServiceState operates on VpcVRouterNetworkServiceState
func (cli *ZSClient) SetVpcVRouterNetworkServiceState(uuid string, params param.SetVpcVRouterNetworkServiceStateParam) (*view.SetVpcVRouterNetworkServiceStateEventView, error) {
	resp := view.SetVpcVRouterNetworkServiceStateEventView{}
	if err := cli.Post(fmt.Sprintf("v1/vpc/virtual-routers/%s/networkservicestate", uuid), params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AddDnsToVpcRouter adds DnsToVpcRouter
func (cli *ZSClient) AddDnsToVpcRouter(uuid string, params param.AddDnsToVpcRouterParam) (*view.VpcRouterVmInventoryView, error) {
	resp := view.VpcRouterVmInventoryView{}
	if err := cli.Post(fmt.Sprintf("v1/vpc/virtual-routers/%s/dns", uuid), params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// QueryVpcRouter queries VpcRouter list
func (cli *ZSClient) QueryVpcRouter(params *param.QueryParam) ([]view.VpcRouterVmInventoryView, error) {
	var resp []view.VpcRouterVmInventoryView
	return resp, cli.List("v1/vpc/virtual-routers", params, &resp)
}

func (cli *ZSClient) GetVpcRouter(uuid string) (*view.VpcRouterVmInventoryView, error) {
	var resp view.VpcRouterVmInventoryView
	if err := cli.Get("v1/vpc/virtual-routers", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageVpcRouter Pagination
func (cli *ZSClient) PageVpcRouter(params *param.QueryParam) ([]view.VpcRouterVmInventoryView, int, error) {
	var vpcRouters []view.VpcRouterVmInventoryView
	total, err := cli.Page("v1/vpc/virtual-routers", params, &vpcRouters)
	return vpcRouters, total, err
}

// UpdateVirtualRouter updates VirtualRouter
func (cli *ZSClient) UpdateVirtualRouter(vmInstanceUuid string, params param.UpdateVirtualRouterParam) (*view.VirtualRouterVmInventoryView, error) {
	resp := view.VirtualRouterVmInventoryView{}
	if err := cli.PutWithRespKey("v1/vm-instances/appliances/virtual-routers", vmInstanceUuid, "", map[string]interface{}{
		"updateVirtualRouter": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetAttachablePublicL3ForVRouter gets AttachablePublicL3ForVRouter by uuid
func (cli *ZSClient) GetAttachablePublicL3ForVRouter(uuid string) (*view.L3NetworkInventoryView, error) {
	var resp view.L3NetworkInventoryView
	if err := cli.GetWithRespKey("v1/vm-instances/appliances/virtual-routers", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SetVpcVRouterDistributedRoutingEnabled operates on VpcVRouterDistributedRoutingEnabled
func (cli *ZSClient) SetVpcVRouterDistributedRoutingEnabled(uuid string, params param.SetVpcVRouterDistributedRoutingEnabledParam) (*view.SetVpcVRouterDistributedRoutingEnabledEventView, error) {
	resp := view.SetVpcVRouterDistributedRoutingEnabledEventView{}
	if err := cli.Post(fmt.Sprintf("v1/vpc/virtual-routers/%s/distributed-routing", uuid), params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ProvisionVirtualRouterConfig operates on ProvisionVirtualRouterConfig
func (cli *ZSClient) ProvisionVirtualRouterConfig(vmInstanceUuid string, params param.ProvisionVirtualRouterConfigParam) (*view.ApplianceVmInventoryView, error) {
	resp := view.ApplianceVmInventoryView{}
	if err := cli.PutWithRespKey("v1/vm-instances/appliances/virtual-routers", vmInstanceUuid, "", map[string]interface{}{
		"provisionVirtualRouterConfig": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateVpcVRouter creates VpcVRouter
func (cli *ZSClient) CreateVpcVRouter(params param.CreateVpcVRouterParam) (*view.VmInstanceInventoryView, error) {
	resp := view.VmInstanceInventoryView{}
	if err := cli.Post("v1/vpc/virtual-routers", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetAvailableVpcL3Network gets AvailableVpcL3Network by uuid
func (cli *ZSClient) GetAvailableVpcL3Network() (*view.GetAvailableVpcL3NetworkView, error) {
	var resp view.GetAvailableVpcL3NetworkView
	if err := cli.GetWithRespKey("v1/vpc/virtual-routers/available-vpc-l3s", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetVpcVRouterDistributedRoutingEnabled gets VpcVRouterDistributedRoutingEnabled by uuid
func (cli *ZSClient) GetVpcVRouterDistributedRoutingEnabled(uuid string) (*view.GetVpcVRouterDistributedRoutingEnabledView, error) {
	var resp view.GetVpcVRouterDistributedRoutingEnabledView
	if err := cli.GetWithRespKey("v1/vpc/virtual-routers", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
