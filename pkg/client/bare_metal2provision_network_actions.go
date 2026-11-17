// Copyright (c) ZStack.io, Inc.

package client

import (
	"fmt"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// UpdateBareMetal2ProvisionNetwork updates BareMetal2ProvisionNetwork
func (cli *ZSClient) UpdateBareMetal2ProvisionNetwork(uuid string, params param.UpdateBareMetal2ProvisionNetworkParam) (*view.BareMetal2ProvisionNetworkInventoryView, error) {
	resp := view.BareMetal2ProvisionNetworkInventoryView{}
	if err := cli.PutWithRespKey("v1/baremetal2/provision-networks", uuid, "", map[string]interface{}{
		"updateBareMetal2ProvisionNetwork": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// QueryBareMetal2ProvisionNetwork queries BareMetal2ProvisionNetwork list
func (cli *ZSClient) QueryBareMetal2ProvisionNetwork(params *param.QueryParam) ([]view.BareMetal2ProvisionNetworkInventoryView, error) {
	var resp []view.BareMetal2ProvisionNetworkInventoryView
	return resp, cli.List("v1/baremetal2/provision-networks", params, &resp)
}

func (cli *ZSClient) GetBareMetal2ProvisionNetwork(uuid string) (*view.BareMetal2ProvisionNetworkInventoryView, error) {
	var resp view.BareMetal2ProvisionNetworkInventoryView
	if err := cli.Get("v1/baremetal2/provision-networks", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageBareMetal2ProvisionNetwork Pagination
func (cli *ZSClient) PageBareMetal2ProvisionNetwork(params *param.QueryParam) ([]view.BareMetal2ProvisionNetworkInventoryView, int, error) {
	var bareMetal2ProvisionNetworks []view.BareMetal2ProvisionNetworkInventoryView
	total, err := cli.Page("v1/baremetal2/provision-networks", params, &bareMetal2ProvisionNetworks)
	return bareMetal2ProvisionNetworks, total, err
}
// CreateBareMetal2ProvisionNetwork creates BareMetal2ProvisionNetwork
func (cli *ZSClient) CreateBareMetal2ProvisionNetwork(params param.CreateBareMetal2ProvisionNetworkParam) (*view.BareMetal2ProvisionNetworkInventoryView, error) {
	resp := view.BareMetal2ProvisionNetworkInventoryView{}
	if err := cli.Post("v1/baremetal2/provision-networks", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// DeleteBareMetal2ProvisionNetwork deletes BareMetal2ProvisionNetwork
func (cli *ZSClient) DeleteBareMetal2ProvisionNetwork(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/baremetal2/provision-networks", uuid, string(deleteMode))
}

// AttachBareMetal2ProvisionNetworkToCluster operates on BareMetal2ProvisionNetworkToCluster
func (cli *ZSClient) AttachBareMetal2ProvisionNetworkToCluster(clusterUuid string, networkUuid string, params param.AttachBareMetal2ProvisionNetworkToClusterParam) (*view.BareMetal2ProvisionNetworkInventoryView, error) {
	resp := view.BareMetal2ProvisionNetworkInventoryView{}
	if err := cli.Post(fmt.Sprintf("v1/baremetal2/clusters/%s/provision-networks/%s", clusterUuid, networkUuid), params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DetachBareMetal2ProvisionNetworkFromCluster operates on BareMetal2ProvisionNetworkFromCluster
func (cli *ZSClient) DetachBareMetal2ProvisionNetworkFromCluster(clusterUuid string, networkUuid string, deleteMode param.DeleteMode) error {
	return cli.DeleteWithSpec("v1/baremetal2/clusters", clusterUuid, fmt.Sprintf("provision-networks/%s", networkUuid), fmt.Sprintf("deleteMode=%s", deleteMode), nil)
}

// GetBareMetal2ProvisionNetworkIpAddressCapacity gets BareMetal2ProvisionNetworkIpAddressCapacity by uuid
func (cli *ZSClient) GetBareMetal2ProvisionNetworkIpAddressCapacity() (*view.GetBareMetal2ProvisionNetworkIpAddressCapacityView, error) {
	var resp view.GetBareMetal2ProvisionNetworkIpAddressCapacityView
	if err := cli.GetWithRespKey("v1/baremetal2/provision-networks/ip-capacity", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ChangeBareMetal2ProvisionNetworkState changes BareMetal2ProvisionNetworkState
func (cli *ZSClient) ChangeBareMetal2ProvisionNetworkState(uuid string, params param.ChangeBareMetal2ProvisionNetworkStateParam) (*view.BareMetal2ProvisionNetworkInventoryView, error) {
	resp := view.BareMetal2ProvisionNetworkInventoryView{}
	if err := cli.PutWithRespKey("v1/baremetal2/provision-networks", uuid, "", map[string]interface{}{
		"changeBareMetal2ProvisionNetworkState": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
