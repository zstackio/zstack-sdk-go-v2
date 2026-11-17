// Copyright (c) ZStack.io, Inc.

package client

import (
	"fmt"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// CreatePortMirror creates PortMirror
func (cli *ZSClient) CreatePortMirror(params param.CreatePortMirrorParam) (*view.PortMirrorInventoryView, error) {
	resp := view.PortMirrorInventoryView{}
	if err := cli.Post("v1/port-mirrors", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// QueryPortMirror queries PortMirror list
func (cli *ZSClient) QueryPortMirror(params *param.QueryParam) ([]view.PortMirrorInventoryView, error) {
	var resp []view.PortMirrorInventoryView
	return resp, cli.List("v1/port-mirrors", params, &resp)
}

func (cli *ZSClient) GetPortMirror(uuid string) (*view.PortMirrorInventoryView, error) {
	var resp view.PortMirrorInventoryView
	if err := cli.Get("v1/port-mirrors", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PagePortMirror Pagination
func (cli *ZSClient) PagePortMirror(params *param.QueryParam) ([]view.PortMirrorInventoryView, int, error) {
	var portMirrors []view.PortMirrorInventoryView
	total, err := cli.Page("v1/port-mirrors", params, &portMirrors)
	return portMirrors, total, err
}
// DeletePortMirror deletes PortMirror
func (cli *ZSClient) DeletePortMirror(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/port-mirrors", uuid, string(deleteMode))
}
// UpdatePortMirror updates PortMirror
func (cli *ZSClient) UpdatePortMirror(uuid string, params param.UpdatePortMirrorParam) (*view.PortMirrorInventoryView, error) {
	resp := view.PortMirrorInventoryView{}
	if err := cli.PutWithRespKey("v1/port-mirrors", uuid, "", map[string]interface{}{
		"updatePortMirror": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ChangePortMirrorState changes PortMirrorState
func (cli *ZSClient) ChangePortMirrorState(uuid string, params param.ChangePortMirrorStateParam) (*view.PortMirrorInventoryView, error) {
	resp := view.PortMirrorInventoryView{}
	if err := cli.PutWithRespKey("v1/port-mirrors", uuid, "", map[string]interface{}{
		"changePortMirrorState": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetCandidateVmNicsForPortMirror gets CandidateVmNicsForPortMirror by uuid
func (cli *ZSClient) GetCandidateVmNicsForPortMirror(portMirrorUuid string, typeParam string) (*view.VmNicInventoryView, error) {
	var resp view.VmNicInventoryView
	err := cli.GetWithSpec("v1/port-mirrors", portMirrorUuid, fmt.Sprintf("vm-instances/candidate-nics/%s", typeParam), "", nil, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// QueryPortMirrorNetworkUsedIp queries PortMirrorNetworkUsedIp list
func (cli *ZSClient) QueryPortMirrorNetworkUsedIp(params *param.QueryParam) ([]view.MirrorNetworkUsedIpInventoryView, error) {
	var resp []view.MirrorNetworkUsedIpInventoryView
	return resp, cli.List("v1/port-mirrors/networks/usedIps", params, &resp)
}

func (cli *ZSClient) GetPortMirrorNetworkUsedIp(uuid string) (*view.MirrorNetworkUsedIpInventoryView, error) {
	var resp view.MirrorNetworkUsedIpInventoryView
	if err := cli.Get("v1/port-mirrors/networks/usedIps", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PagePortMirrorNetworkUsedIp Pagination
func (cli *ZSClient) PagePortMirrorNetworkUsedIp(params *param.QueryParam) ([]view.MirrorNetworkUsedIpInventoryView, int, error) {
	var portMirrorNetworkUsedIps []view.MirrorNetworkUsedIpInventoryView
	total, err := cli.Page("v1/port-mirrors/networks/usedIps", params, &portMirrorNetworkUsedIps)
	return portMirrorNetworkUsedIps, total, err
}
