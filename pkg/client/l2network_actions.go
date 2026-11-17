// Copyright (c) ZStack.io, Inc.

package client

import (
	"fmt"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// DeleteL2Network deletes L2Network
func (cli *ZSClient) DeleteL2Network(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/l2-networks", uuid, string(deleteMode))
}
// UpdateL2Network updates L2Network
func (cli *ZSClient) UpdateL2Network(uuid string, params param.UpdateL2NetworkParam) (*view.L2NetworkInventoryView, error) {
	resp := view.L2NetworkInventoryView{}
	if err := cli.PutWithRespKey("v1/l2-networks", uuid, "", map[string]interface{}{
		"updateL2Network": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// QueryL2Network queries L2Network list
func (cli *ZSClient) QueryL2Network(params *param.QueryParam) ([]view.L2NetworkInventoryView, error) {
	var resp []view.L2NetworkInventoryView
	return resp, cli.List("v1/l2-networks", params, &resp)
}

func (cli *ZSClient) GetL2Network(uuid string) (*view.L2NetworkInventoryView, error) {
	var resp view.L2NetworkInventoryView
	if err := cli.Get("v1/l2-networks", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageL2Network Pagination
func (cli *ZSClient) PageL2Network(params *param.QueryParam) ([]view.L2NetworkInventoryView, int, error) {
	var l2Networks []view.L2NetworkInventoryView
	total, err := cli.Page("v1/l2-networks", params, &l2Networks)
	return l2Networks, total, err
}

// AttachL2NetworkToCluster operates on L2NetworkToCluster
func (cli *ZSClient) AttachL2NetworkToCluster(l2NetworkUuid string, clusterUuid string, params param.AttachL2NetworkToClusterParam) (*view.L2NetworkInventoryView, error) {
	resp := view.L2NetworkInventoryView{}
	body := struct {
		param.BaseParam
		L2ProviderType *string `json:"l2ProviderType,omitempty"`
	}{
		BaseParam:      params.BaseParam,
		L2ProviderType: params.Params.L2ProviderType,
	}
	if err := cli.Post(fmt.Sprintf("v1/l2-networks/%s/clusters/%s", l2NetworkUuid, clusterUuid), body, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetCandidateClustersForAttachingL2Network gets CandidateClustersForAttachingL2Network by uuid
func (cli *ZSClient) GetCandidateClustersForAttachingL2Network(uuid string) (*view.ClusterInventoryView, error) {
	var resp view.ClusterInventoryView
	if err := cli.GetWithRespKey("v1/l2-networks", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateL2TfNetwork creates L2TfNetwork
func (cli *ZSClient) CreateL2TfNetwork(params param.CreateL2TfNetworkParam) (*view.L2NetworkInventoryView, error) {
	resp := view.L2NetworkInventoryView{}
	if err := cli.Post("v1/l2-networks/tf", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetCandidateL2NetworksForAttachingCluster gets CandidateL2NetworksForAttachingCluster by uuid
func (cli *ZSClient) GetCandidateL2NetworksForAttachingCluster(uuid string) (*view.L2NetworkDataView, error) {
	var resp view.L2NetworkDataView
	if err := cli.GetWithRespKey("v1/cluster", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateL2NoVlanNetwork creates L2NoVlanNetwork
func (cli *ZSClient) CreateL2NoVlanNetwork() (*view.L2NetworkInventoryView, error) {
	resp := view.L2NetworkInventoryView{}
	if err := cli.Post("v1/l2-networks/no-vlan", map[string]interface{}{}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetVSwitchTypes gets VSwitchTypes by uuid
func (cli *ZSClient) GetVSwitchTypes() (*view.GetVSwitchTypesView, error) {
	var resp view.GetVSwitchTypesView
	if err := cli.GetWithRespKey("v1/l2-networks/vSwitchTypes", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DetachL2NetworkFromCluster operates on L2NetworkFromCluster
func (cli *ZSClient) DetachL2NetworkFromCluster(l2NetworkUuid string, clusterUuid string, deleteMode param.DeleteMode) error {
	return cli.DeleteWithSpec("v1/l2-networks", l2NetworkUuid, fmt.Sprintf("clusters/%s", clusterUuid), fmt.Sprintf("deleteMode=%s", deleteMode), nil)
}

// GetVfPciDeviceAvailableInL2Network gets VfPciDeviceAvailableInL2Network by uuid
func (cli *ZSClient) GetVfPciDeviceAvailableInL2Network() (*view.GetVfPciDeviceAvailableInL2NetworkView, error) {
	var resp view.GetVfPciDeviceAvailableInL2NetworkView
	if err := cli.GetWithRespKey("v1/l2-networks/vf-pci-devices-available", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetL2NetworkTypes gets L2NetworkTypes by uuid
func (cli *ZSClient) GetL2NetworkTypes() (*view.GetL2NetworkTypesView, error) {
	var resp view.GetL2NetworkTypesView
	if err := cli.GetWithRespKey("v1/l2-networks/types", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
