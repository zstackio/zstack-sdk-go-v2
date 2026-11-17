// Copyright (c) ZStack.io, Inc.

package client

import (
	"fmt"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// ReconnectBareMetal2Gateway operates on BareMetal2Gateway
func (cli *ZSClient) ReconnectBareMetal2Gateway(uuid string, params param.ReconnectBareMetal2GatewayParam) (*view.BareMetal2GatewayInventoryView, error) {
	resp := view.BareMetal2GatewayInventoryView{}
	if err := cli.PutWithRespKey("v1/baremetal2/gateways", uuid, "", map[string]interface{}{
		"reconnectBareMetal2Gateway": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// UpdateBareMetal2Gateway updates BareMetal2Gateway
func (cli *ZSClient) UpdateBareMetal2Gateway(uuid string, params param.UpdateBareMetal2GatewayParam) (*view.BareMetal2GatewayInventoryView, error) {
	resp := view.BareMetal2GatewayInventoryView{}
	if err := cli.PutWithRespKey("v1/baremetal2/gateways", uuid, "", map[string]interface{}{
		"updateBareMetal2Gateway": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// QueryBareMetal2Gateway queries BareMetal2Gateway list
func (cli *ZSClient) QueryBareMetal2Gateway(params *param.QueryParam) ([]view.BareMetal2GatewayInventoryView, error) {
	var resp []view.BareMetal2GatewayInventoryView
	return resp, cli.List("v1/baremetal2/gateways", params, &resp)
}

func (cli *ZSClient) GetBareMetal2Gateway(uuid string) (*view.BareMetal2GatewayInventoryView, error) {
	var resp view.BareMetal2GatewayInventoryView
	if err := cli.Get("v1/baremetal2/gateways", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageBareMetal2Gateway Pagination
func (cli *ZSClient) PageBareMetal2Gateway(params *param.QueryParam) ([]view.BareMetal2GatewayInventoryView, int, error) {
	var bareMetal2Gatewaies []view.BareMetal2GatewayInventoryView
	total, err := cli.Page("v1/baremetal2/gateways", params, &bareMetal2Gatewaies)
	return bareMetal2Gatewaies, total, err
}
// DeleteBareMetal2Gateway deletes BareMetal2Gateway
func (cli *ZSClient) DeleteBareMetal2Gateway(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/baremetal2/gateways", uuid, string(deleteMode))
}
// AddBareMetal2Gateway adds BareMetal2Gateway
func (cli *ZSClient) AddBareMetal2Gateway(params param.AddBareMetal2GatewayParam) (*view.HostInventoryView, error) {
	resp := view.HostInventoryView{}
	if err := cli.Post("v1/baremetal2/gateways", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DetachBareMetal2GatewayFromCluster operates on BareMetal2GatewayFromCluster
func (cli *ZSClient) DetachBareMetal2GatewayFromCluster(clusterUuid string, gatewayUuid string, deleteMode param.DeleteMode) error {
	return cli.DeleteWithSpec("v1/baremetal2/clusters", clusterUuid, fmt.Sprintf("gateways/%s", gatewayUuid), fmt.Sprintf("deleteMode=%s", deleteMode), nil)
}

// ChangeBareMetal2GatewayCluster changes BareMetal2GatewayCluster
func (cli *ZSClient) ChangeBareMetal2GatewayCluster(gatewayUuid string, params param.ChangeBareMetal2GatewayClusterParam) (*view.BareMetal2GatewayInventoryView, error) {
	resp := view.BareMetal2GatewayInventoryView{}
	if err := cli.PutWithRespKey("v1/baremetal2/gateways", gatewayUuid, "", map[string]interface{}{
		"changeBareMetal2GatewayCluster": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ChangeBareMetal2GatewayState changes BareMetal2GatewayState
func (cli *ZSClient) ChangeBareMetal2GatewayState(uuid string, params param.ChangeBareMetal2GatewayStateParam) (*view.BareMetal2GatewayInventoryView, error) {
	resp := view.BareMetal2GatewayInventoryView{}
	if err := cli.PutWithRespKey("v1/baremetal2/gateways", uuid, "", map[string]interface{}{
		"changeBareMetal2GatewayState": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetBareMetal2GatewayAllocatorStrategies gets BareMetal2GatewayAllocatorStrategies by uuid
func (cli *ZSClient) GetBareMetal2GatewayAllocatorStrategies() (*view.GetBareMetal2GatewayAllocatorStrategiesView, error) {
	var resp view.GetBareMetal2GatewayAllocatorStrategiesView
	if err := cli.GetWithRespKey("v1/baremetal2/gateways/allocators/strategies", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AttachBareMetal2GatewayToCluster operates on BareMetal2GatewayToCluster
func (cli *ZSClient) AttachBareMetal2GatewayToCluster(clusterUuid string, gatewayUuid string, params param.AttachBareMetal2GatewayToClusterParam) (*view.BareMetal2GatewayInventoryView, error) {
	resp := view.BareMetal2GatewayInventoryView{}
	if err := cli.Post(fmt.Sprintf("v1/baremetal2/clusters/%s/gateways/%s", clusterUuid, gatewayUuid), params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
