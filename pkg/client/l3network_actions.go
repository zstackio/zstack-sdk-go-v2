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

// QueryL3Network queries L3Network list
func (cli *ZSClient) QueryL3Network(params *param.QueryParam) ([]view.L3NetworkInventoryView, error) {
	var resp []view.L3NetworkInventoryView
	return resp, cli.List("v1/l3-networks", params, &resp)
}

func (cli *ZSClient) GetL3Network(uuid string) (*view.L3NetworkInventoryView, error) {
	var resp view.L3NetworkInventoryView
	if err := cli.Get("v1/l3-networks", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageL3Network Pagination
func (cli *ZSClient) PageL3Network(params *param.QueryParam) ([]view.L3NetworkInventoryView, int, error) {
	var l3Networks []view.L3NetworkInventoryView
	total, err := cli.Page("v1/l3-networks", params, &l3Networks)
	return l3Networks, total, err
}
// UpdateL3Network updates L3Network
func (cli *ZSClient) UpdateL3Network(uuid string, params param.UpdateL3NetworkParam) (*view.L3NetworkInventoryView, error) {
	resp := view.L3NetworkInventoryView{}
	if err := cli.PutWithRespKey("v1/l3-networks", uuid, "", map[string]interface{}{
		"updateL3Network": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// CreateL3Network creates L3Network
func (cli *ZSClient) CreateL3Network(params param.CreateL3NetworkParam) (*view.L3NetworkInventoryView, error) {
	resp := view.L3NetworkInventoryView{}
	if err := cli.Post("v1/l3-networks", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// DeleteL3Network deletes L3Network
func (cli *ZSClient) DeleteL3Network(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/l3-networks", uuid, string(deleteMode))
}

// GetInterdependentL3NetworksBackupStorages gets InterdependentL3NetworksBackupStorages by uuid
func (cli *ZSClient) GetInterdependentL3NetworksBackupStorages() (*view.GetInterdependentL3NetworksBackupStoragesView, error) {
	var resp view.GetInterdependentL3NetworksBackupStoragesView
	if err := cli.GetWithRespKey("v1/backupStorage-l3networks/dependencies", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ChangeL3NetworkDhcpIpAddress changes L3NetworkDhcpIpAddress
func (cli *ZSClient) ChangeL3NetworkDhcpIpAddress(l3NetworkUuid string, params param.ChangeL3NetworkDhcpIpAddressParam) (*view.ChangeL3NetworkDhcpIpAddressEventView, error) {
	resp := view.ChangeL3NetworkDhcpIpAddressEventView{}
	if err := cli.PutWithRespKey("v1/l3-networks", l3NetworkUuid, "", map[string]interface{}{
		"changeL3NetworkDhcpIpAddress": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetInterdependentL3NetworksImages gets InterdependentL3NetworksImages by uuid
func (cli *ZSClient) GetInterdependentL3NetworksImages() (*view.GetInterdependentL3NetworkImageView, error) {
	var resp view.GetInterdependentL3NetworkImageView
	if err := cli.GetWithRespKey("v1/images-l3networks/dependencies", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// RemoveHostRouteFromL3Network requires additional arguments.
//
// Deprecated: use RemoveHostRouteFromL3NetworkWithParams to supply the required fields.
func (cli *ZSClient) RemoveHostRouteFromL3Network(uuid string, deleteMode param.DeleteMode) error {
	return fmt.Errorf("RemoveHostRouteFromL3Network requires additional arguments; use RemoveHostRouteFromL3NetworkWithParams")
}

// RemoveHostRouteFromL3NetworkWithParams sends the required route and query parameters for RemoveHostRouteFromL3Network.
func (cli *ZSClient) RemoveHostRouteFromL3NetworkWithParams(l3NetworkUuid string, deleteMode param.DeleteMode, params param.RemoveHostRouteFromL3NetworkParamDetail) error {
	query := url.Values{
		"prefix": {params.Prefix},
	}
	return cli.deleteActionWithQuery("v1/l3-networks", l3NetworkUuid, "hostroute", deleteMode, query, "prefix")
}

// IsVfNicAvailableInL3Network operates on IsVfNicAvailableInL3Network
func (cli *ZSClient) IsVfNicAvailableInL3Network(l3NetworkUuid string, hostUuid string) (*view.IsVfNicAvailableInL3NetworkView, error) {
	var resp view.IsVfNicAvailableInL3NetworkView
	err := cli.GetWithSpec("v1/l3-networks", l3NetworkUuid, fmt.Sprintf("hosts/%s/vfnicavailable", hostUuid), "", nil, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetCandidateL3NetworksForServerGroup gets CandidateL3NetworksForServerGroup by uuid
func (cli *ZSClient) GetCandidateL3NetworksForServerGroup() (*view.L3NetworkInventoryView, error) {
	var resp view.L3NetworkInventoryView
	if err := cli.GetWithRespKey("v1/load-balancers/servergroups/candidate-l3network", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AddHostRouteToL3Network adds HostRouteToL3Network
func (cli *ZSClient) AddHostRouteToL3Network(l3NetworkUuid string, params param.AddHostRouteToL3NetworkParam) (*view.L3NetworkInventoryView, error) {
	resp := view.L3NetworkInventoryView{}
	if err := cli.Post(fmt.Sprintf("v1/l3-networks/%s/hostroute", l3NetworkUuid), params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetL3NetworkIpStatistic gets L3NetworkIpStatistic by uuid
func (cli *ZSClient) GetL3NetworkIpStatistic(uuid string) (*view.GetL3NetworkIpStatisticView, error) {
	var resp view.GetL3NetworkIpStatisticView
	if err := cli.GetWithRespKey("v1/l3-networks", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CheckIpAvailability operates on IpAvailability
func (cli *ZSClient) CheckIpAvailability(l3NetworkUuid string, ip string) (*view.CheckIpAvailabilityView, error) {
	var resp view.CheckIpAvailabilityView
	err := cli.GetWithSpec("v1/l3-networks", l3NetworkUuid, fmt.Sprintf("ip/%s/availability", ip), "", nil, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetIpAddressCapacity gets IpAddressCapacity by uuid
func (cli *ZSClient) GetIpAddressCapacity() (*view.GetIpAddressCapacityView, error) {
	var resp view.GetIpAddressCapacityView
	if err := cli.GetWithRespKey("v1/ip-capacity", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// RemoveDnsFromL3Network removes DnsFromL3Network
func (cli *ZSClient) RemoveDnsFromL3Network(l3NetworkUuid string, dns string, deleteMode param.DeleteMode) error {
	return cli.DeleteWithSpec("v1/l3-networks", l3NetworkUuid, fmt.Sprintf("dns/%s", dns), fmt.Sprintf("deleteMode=%s", deleteMode), nil)
}

// AddL3NetworkToGroup adds L3NetworkToGroup
func (cli *ZSClient) AddL3NetworkToGroup(nfvInstGroupUuid string, networkServiceUuid string, params param.AddL3NetworkToGroupParam) (*view.AddL3NetworkToGroupEventView, error) {
	resp := view.AddL3NetworkToGroupEventView{}
	if err := cli.Post(fmt.Sprintf("v1/nfvinstgroup/group/%s/service/%s", nfvInstGroupUuid, networkServiceUuid), params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ChangeL3NetworkState changes L3NetworkState
func (cli *ZSClient) ChangeL3NetworkState(uuid string, params param.ChangeL3NetworkStateParam) (*view.L3NetworkInventoryView, error) {
	resp := view.L3NetworkInventoryView{}
	if err := cli.PutWithRespKey("v1/l3-networks", uuid, "", map[string]interface{}{
		"changeL3NetworkState": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetL3NetworkDhcpIpAddress gets L3NetworkDhcpIpAddress by uuid
func (cli *ZSClient) GetL3NetworkDhcpIpAddress(uuid string) (*view.GetL3NetworkDhcpIpAddressView, error) {
	var resp view.GetL3NetworkDhcpIpAddressView
	if err := cli.GetWithRespKey("v1/l3-networks", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetAttachableVpcL3Network gets AttachableVpcL3Network by uuid
func (cli *ZSClient) GetAttachableVpcL3Network(uuid string, params param.GetAttachableVpcL3NetworkParam) (*view.L3NetworkInventoryView, error) {
	resp := view.L3NetworkInventoryView{}
	if err := cli.Post(fmt.Sprintf("v1/vpc/virtual-routers/%s/attachable-vpc-l3s", uuid), params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetL3NetworkMtu gets L3NetworkMtu by uuid
func (cli *ZSClient) GetL3NetworkMtu(uuid string) (*view.GetL3NetworkMtuView, error) {
	var resp view.GetL3NetworkMtuView
	if err := cli.GetWithRespKey("v1/l3-networks", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AddDnsToL3Network adds DnsToL3Network
func (cli *ZSClient) AddDnsToL3Network(l3NetworkUuid string, params param.AddDnsToL3NetworkParam) (*view.L3NetworkInventoryView, error) {
	resp := view.L3NetworkInventoryView{}
	if err := cli.Post(fmt.Sprintf("v1/l3-networks/%s/dns", l3NetworkUuid), params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetCandidateL3NetworksForIpSecConnection gets CandidateL3NetworksForIpSecConnection by uuid
func (cli *ZSClient) GetCandidateL3NetworksForIpSecConnection() (*view.L3NetworkInventoryView, error) {
	var resp view.L3NetworkInventoryView
	if err := cli.GetWithRespKey("v1/ipsec/candidatesL3Networks", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetVmAttachableL3Network gets VmAttachableL3Network by uuid
func (cli *ZSClient) GetVmAttachableL3Network(uuid string) (*view.L3NetworkInventoryView, error) {
	var resp view.L3NetworkInventoryView
	if err := cli.GetWithRespKey("v1/vm-instances", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SetL3NetworkMtu operates on L3NetworkMtu
func (cli *ZSClient) SetL3NetworkMtu(l3NetworkUuid string, params param.SetL3NetworkMtuParam) (*view.SetL3NetworkMtuEventView, error) {
	resp := view.SetL3NetworkMtuEventView{}
	if err := cli.Post(fmt.Sprintf("v1/l3-networks/%s/mtu", l3NetworkUuid), params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetL3NetworkRouterInterfaceIp gets L3NetworkRouterInterfaceIp by uuid
func (cli *ZSClient) GetL3NetworkRouterInterfaceIp(uuid string) (*view.GetL3NetworkRouterInterfaceIpView, error) {
	var resp view.GetL3NetworkRouterInterfaceIpView
	if err := cli.GetWithRespKey("v1/l3-networks", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SetL3NetworkRouterInterfaceIp operates on L3NetworkRouterInterfaceIp
func (cli *ZSClient) SetL3NetworkRouterInterfaceIp(l3NetworkUuid string, params param.SetL3NetworkRouterInterfaceIpParam) (*view.SetL3NetworkRouterInterfaceIpEventView, error) {
	resp := view.SetL3NetworkRouterInterfaceIpEventView{}
	if err := cli.Post(fmt.Sprintf("v1/l3-networks/%s/router-interface-ip", l3NetworkUuid), params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetCandidateL3NetworksForLoadBalancer gets CandidateL3NetworksForLoadBalancer by uuid
func (cli *ZSClient) GetCandidateL3NetworksForLoadBalancer(uuid string) (*view.L3NetworkInventoryView, error) {
	var resp view.L3NetworkInventoryView
	if err := cli.GetWithRespKey("v1/load-balancers/listeners", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetL3NetworkTypes gets L3NetworkTypes by uuid
func (cli *ZSClient) GetL3NetworkTypes() (*view.GetL3NetworkTypesView, error) {
	var resp view.GetL3NetworkTypesView
	if err := cli.GetWithRespKey("v1/l3-networks/types", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
