// Copyright (c) ZStack.io, Inc.

package client

import (
	"fmt"
	"net/url"

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

// DeleteVxlanPoolRemoteVtep requires the VTEP address.
//
// Deprecated: use DeleteVxlanPoolRemoteVtepWithParams to specify the address.
func (cli *ZSClient) DeleteVxlanPoolRemoteVtep(l2NetworkUuid string, clusterUuid string, deleteMode param.DeleteMode) error {
	return fmt.Errorf("DeleteVxlanPoolRemoteVtep requires remoteVtepIp; use DeleteVxlanPoolRemoteVtepWithParams")
}

// DeleteVxlanPoolRemoteVtepWithParams deletes one remote VTEP address.
func (cli *ZSClient) DeleteVxlanPoolRemoteVtepWithParams(l2NetworkUuid, clusterUuid string, deleteMode param.DeleteMode, params param.DeleteVxlanPoolRemoteVtepParamDetail) error {
	cluster, err := actionPathSegment("clusterUuid", clusterUuid)
	if err != nil {
		return err
	}
	query := url.Values{"remoteVtepIp": {params.RemoteVtepIp}}
	return cli.deleteActionWithQuery("v1/l2-networks", l2NetworkUuid, "clusters/"+cluster+"/delete/remote-vtep-ip",
		deleteModeOrDefault(deleteMode, params.DeleteMode), query, "remoteVtepIp")
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
