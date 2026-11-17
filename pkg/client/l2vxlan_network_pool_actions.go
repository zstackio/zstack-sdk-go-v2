// Copyright (c) ZStack.io, Inc.

package client

import (
	"fmt"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{}    // avoid unused import

// CreateL2VxlanNetworkPool creates L2VxlanNetworkPool
func (cli *ZSClient) CreateL2VxlanNetworkPool(params param.CreateL2VxlanNetworkPoolParam) (*view.L2VxlanNetworkPoolInventoryView, error) {
	resp := view.L2VxlanNetworkPoolInventoryView{}
	if err := cli.Post("v1/l2-networks/vxlan-pool", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// QueryL2VxlanNetworkPool queries L2VxlanNetworkPool list
func (cli *ZSClient) QueryL2VxlanNetworkPool(params *param.QueryParam) ([]view.L2VxlanNetworkPoolInventoryView, error) {
	var resp []view.L2VxlanNetworkPoolInventoryView
	return resp, cli.List("v1/l2-networks/vxlan-pool", params, &resp)
}

func (cli *ZSClient) GetL2VxlanNetworkPool(uuid string) (*view.L2VxlanNetworkPoolInventoryView, error) {
	var resp view.L2VxlanNetworkPoolInventoryView
	if err := cli.Get("v1/l2-networks/vxlan-pool", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageL2VxlanNetworkPool Pagination
func (cli *ZSClient) PageL2VxlanNetworkPool(params *param.QueryParam) ([]view.L2VxlanNetworkPoolInventoryView, int, error) {
	var l2VxlanNetworkPools []view.L2VxlanNetworkPoolInventoryView
	total, err := cli.Page("v1/l2-networks/vxlan-pool", params, &l2VxlanNetworkPools)
	return l2VxlanNetworkPools, total, err
}

// DeleteVxlanPoolRemoteVtep deletes VxlanPoolRemoteVtep
func (cli *ZSClient) DeleteVxlanPoolRemoteVtep(l2NetworkUuid string, clusterUuid string, deleteMode param.DeleteMode) error {
	return cli.DeleteWithSpec("v1/l2-networks", l2NetworkUuid, fmt.Sprintf("clusters/%s/delete/remote-vtep-ip", clusterUuid), fmt.Sprintf("deleteMode=%s", deleteMode), nil)
}

// CreateL2HardwareVxlanNetworkPool creates L2HardwareVxlanNetworkPool
func (cli *ZSClient) CreateL2HardwareVxlanNetworkPool(params param.CreateL2HardwareVxlanNetworkPoolParam) (*view.CreateL2HardwareVxlanNetworkPoolEventView, error) {
	resp := view.CreateL2HardwareVxlanNetworkPoolEventView{}
	if err := cli.Post("v1/l2-networks/hardware-vxlan-pool", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateVxlanPoolRemoteVtep creates VxlanPoolRemoteVtep
func (cli *ZSClient) CreateVxlanPoolRemoteVtep(l2NetworkUuid string, clusterUuid string, params param.CreateVxlanPoolRemoteVtepParam) (*view.RemoteVtepInventoryView, error) {
	resp := view.RemoteVtepInventoryView{}
	if err := cli.Post(fmt.Sprintf("v1/l2-networks/%s/clusters/%s/remote-vtep-ip", l2NetworkUuid, clusterUuid), params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
