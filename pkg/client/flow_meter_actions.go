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

// QueryFlowMeter queries FlowMeter list
func (cli *ZSClient) QueryFlowMeter(params *param.QueryParam) ([]view.FlowMeterInventoryView, error) {
	var resp []view.FlowMeterInventoryView
	return resp, cli.List("v1/flowmeters", params, &resp)
}

func (cli *ZSClient) GetFlowMeter(uuid string) (*view.FlowMeterInventoryView, error) {
	var resp view.FlowMeterInventoryView
	if err := cli.Get("v1/flowmeters", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageFlowMeter Pagination
func (cli *ZSClient) PageFlowMeter(params *param.QueryParam) ([]view.FlowMeterInventoryView, int, error) {
	var flowMeters []view.FlowMeterInventoryView
	total, err := cli.Page("v1/flowmeters", params, &flowMeters)
	return flowMeters, total, err
}
// CreateFlowMeter creates FlowMeter
func (cli *ZSClient) CreateFlowMeter(params param.CreateFlowMeterParam) (*view.FlowMeterInventoryView, error) {
	resp := view.FlowMeterInventoryView{}
	if err := cli.Post("v1/flowmeters", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// DeleteFlowMeter deletes FlowMeter
func (cli *ZSClient) DeleteFlowMeter(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/flowmeters", uuid, string(deleteMode))
}
// UpdateFlowMeter updates FlowMeter
func (cli *ZSClient) UpdateFlowMeter(uuid string, params param.UpdateFlowMeterParam) (*view.FlowMeterInventoryView, error) {
	resp := view.FlowMeterInventoryView{}
	if err := cli.PutWithRespKey("v1/flowmeters", uuid, "", map[string]interface{}{
		"updateFlowMeter": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetVRouterFlowCounter gets VRouterFlowCounter by uuid
func (cli *ZSClient) GetVRouterFlowCounter(uuid string) (*view.GetVRouterFlowCounterView, error) {
	var resp view.GetVRouterFlowCounterView
	if err := cli.GetWithRespKey("v1/flowmeters", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetVpcAttachedNetflow gets VpcAttachedNetflow by uuid
func (cli *ZSClient) GetVpcAttachedNetflow(uuid string, params param.GetVpcAttachedNetflowParam) (*view.FlowMeterInventoryView, error) {
	resp := view.FlowMeterInventoryView{}
	if err := cli.Post(fmt.Sprintf("v1/vpc/virtual-routers/%s/attached-netflow", uuid), params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// QueryVRouterFlowMeterNetwork queries VRouterFlowMeterNetwork list
func (cli *ZSClient) QueryVRouterFlowMeterNetwork(params *param.QueryParam) ([]view.NetworkRouterFlowMeterRefInventoryView, error) {
	var resp []view.NetworkRouterFlowMeterRefInventoryView
	return resp, cli.List("v1/flowmeters/networks", params, &resp)
}

func (cli *ZSClient) GetVRouterFlowMeterNetwork(uuid string) (*view.NetworkRouterFlowMeterRefInventoryView, error) {
	var resp view.NetworkRouterFlowMeterRefInventoryView
	if err := cli.Get("v1/flowmeters/networks", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageVRouterFlowMeterNetwork Pagination
func (cli *ZSClient) PageVRouterFlowMeterNetwork(params *param.QueryParam) ([]view.NetworkRouterFlowMeterRefInventoryView, int, error) {
	var vRouterFlowMeterNetworks []view.NetworkRouterFlowMeterRefInventoryView
	total, err := cli.Page("v1/flowmeters/networks", params, &vRouterFlowMeterNetworks)
	return vRouterFlowMeterNetworks, total, err
}

// AddVRouterNetworksToFlowMeter adds VRouterNetworksToFlowMeter
func (cli *ZSClient) AddVRouterNetworksToFlowMeter(flowMeterUuid string, vRouterUuid string, params param.AddVRouterNetworksToFlowMeterParam) (*view.NetworkRouterFlowMeterRefInventoryView, error) {
	resp := view.NetworkRouterFlowMeterRefInventoryView{}
	if err := cli.Post(fmt.Sprintf("v1/flowmeters/%s/router/%s/addnetworks", flowMeterUuid, vRouterUuid), params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// RemoveVRouterNetworksFromFlowMeter sends one UUID to the collection endpoint.
func (cli *ZSClient) RemoveVRouterNetworksFromFlowMeter(uuid string, deleteMode param.DeleteMode) error {
	return cli.RemoveVRouterNetworksFromFlowMeterWithParams(deleteMode, param.RemoveVRouterNetworksFromFlowMeterParamDetail{Uuids: []string{uuid}})
}

// RemoveVRouterNetworksFromFlowMeterWithParams sends the DELETE request with its operation parameters.
func (cli *ZSClient) RemoveVRouterNetworksFromFlowMeterWithParams(deleteMode param.DeleteMode, params param.RemoveVRouterNetworksFromFlowMeterParamDetail) error {
	query := url.Values{
		"uuids": params.Uuids,
	}
	return cli.deleteActionWithQuery("v1/flowmeters/networks", "", "", deleteModeOrDefault(deleteMode, params.DeleteMode), query, "uuids")
}

// SetFlowMeterRouterId operates on FlowMeterRouterId
func (cli *ZSClient) SetFlowMeterRouterId(vRouterUuid string, params param.SetFlowMeterRouterIdParam) (*view.SetFlowMeterRouterIdEventView, error) {
	resp := view.SetFlowMeterRouterIdEventView{}
	if err := cli.Post(fmt.Sprintf("v1/flowmeters/%s/routerid", vRouterUuid), params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetFlowMeterRouterId gets FlowMeterRouterId by uuid
func (cli *ZSClient) GetFlowMeterRouterId(uuid string) (*view.GetFlowMeterRouterIdView, error) {
	var resp view.GetFlowMeterRouterIdView
	if err := cli.GetWithRespKey("v1/flowmeters", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
