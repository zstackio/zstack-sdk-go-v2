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

// CreateVmNic creates VmNic
func (cli *ZSClient) CreateVmNic(params param.CreateVmNicParam) (*view.VmNicInventoryView, error) {
	resp := view.VmNicInventoryView{}
	if err := cli.Post("v1/nics", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// QueryVmNic queries VmNic list
func (cli *ZSClient) QueryVmNic(params *param.QueryParam) ([]view.VmNicInventoryView, error) {
	var resp []view.VmNicInventoryView
	return resp, cli.List("v1/vm-instances/nics", params, &resp)
}

func (cli *ZSClient) GetVmNic(uuid string) (*view.VmNicInventoryView, error) {
	var resp view.VmNicInventoryView
	if err := cli.Get("v1/vm-instances/nics", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageVmNic Pagination
func (cli *ZSClient) PageVmNic(params *param.QueryParam) ([]view.VmNicInventoryView, int, error) {
	var vmNics []view.VmNicInventoryView
	total, err := cli.Page("v1/vm-instances/nics", params, &vmNics)
	return vmNics, total, err
}
// DeleteVmNic deletes VmNic
func (cli *ZSClient) DeleteVmNic(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/nics", uuid, string(deleteMode))
}

// ChangeVmNicType changes VmNicType
func (cli *ZSClient) ChangeVmNicType(vmNicUuid string, params param.ChangeVmNicTypeParam) (*view.VmNicInventoryView, error) {
	resp := view.VmNicInventoryView{}
	if err := cli.PutWithRespKey("v1/vm-instances/nics", vmNicUuid, "", map[string]interface{}{
		"changeVmNicType": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateVmNicMac updates VmNicMac
func (cli *ZSClient) UpdateVmNicMac(vmNicUuid string, params param.UpdateVmNicMacParam) (*view.VmNicInventoryView, error) {
	resp := view.VmNicInventoryView{}
	if err := cli.PutWithRespKey("v1/vm-instances/nics", vmNicUuid, "", map[string]interface{}{
		"updateVmNicMac": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteNicQos requires additional arguments.
//
// Deprecated: use DeleteNicQosWithParams to supply the required fields.
func (cli *ZSClient) DeleteNicQos(uuid string, deleteMode param.DeleteMode) error {
	return fmt.Errorf("DeleteNicQos requires additional arguments; use DeleteNicQosWithParams")
}

// DeleteNicQosWithParams sends the required route and query parameters for DeleteNicQos.
func (cli *ZSClient) DeleteNicQosWithParams(uuid string, deleteMode param.DeleteMode, params param.DeleteNicQosParamDetail) error {
	query := url.Values{
		"direction": {params.Direction},
	}
	return cli.deleteActionWithQuery("v1/vm-instances", uuid, "nic-qos", deleteMode, query, "direction")
}

// ChangeVfNicHaState changes VfNicHaState
func (cli *ZSClient) ChangeVfNicHaState(vfNicUuid string, params param.ChangeVfNicHaStateParam) (*view.VmVfNicInventoryView, error) {
	resp := view.VmVfNicInventoryView{}
	if err := cli.PutWithRespKey("v1/vm-instances/nics", vfNicUuid, "", map[string]interface{}{
		"changeVfNicHaState": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SetNicQos operates on NicQos
func (cli *ZSClient) SetNicQos(uuid string, params param.SetNicQosParam) (*view.SetNicQosEventView, error) {
	resp := view.SetNicQosEventView{}
	if err := cli.PutWithRespKey("v1/vm-instances", uuid, "", map[string]interface{}{
		"setNicQos": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetVmNicAttachedNetworkService gets VmNicAttachedNetworkService by uuid
func (cli *ZSClient) GetVmNicAttachedNetworkService(uuid string) (*view.GetVmNicAttachedNetworkServiceView, error) {
	var resp view.GetVmNicAttachedNetworkServiceView
	if err := cli.GetWithRespKey("v1/vm-instances/nics", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DetachL3NetworkFromVm operates on L3NetworkFromVm
func (cli *ZSClient) DetachL3NetworkFromVm(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/vm-instances/nics", uuid, string(deleteMode))
}

// GetPortForwardingAttachableVmNics gets PortForwardingAttachableVmNics by uuid
func (cli *ZSClient) GetPortForwardingAttachableVmNics(uuid string) (*view.VmNicInventoryView, error) {
	var resp view.VmNicInventoryView
	if err := cli.GetWithRespKey("v1/port-forwarding", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetVmNicAttachableEips gets VmNicAttachableEips by uuid
func (cli *ZSClient) GetVmNicAttachableEips(uuid string) (*view.EipInventoryView, error) {
	var resp view.EipInventoryView
	if err := cli.GetWithRespKey("v1/vm-instances/nics", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetCandidateL3NetworksForChangeVmNicNetwork gets CandidateL3NetworksForChangeVmNicNetwork by uuid
func (cli *ZSClient) GetCandidateL3NetworksForChangeVmNicNetwork(uuid string) (*view.L3NetworkInventoryView, error) {
	var resp view.L3NetworkInventoryView
	if err := cli.GetWithRespKey("v1/vm-instances/nics", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetEipAttachableVmNics gets EipAttachableVmNics by uuid
func (cli *ZSClient) GetEipAttachableVmNics(uuid string) (*view.VmNicInventoryView, error) {
	var resp view.VmNicInventoryView
	if err := cli.GetWithRespKey("v1/eips", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AttachL3NetworkToVmNic operates on L3NetworkToVmNic
func (cli *ZSClient) AttachL3NetworkToVmNic(vmNicUuid string, l3NetworkUuid string, params param.AttachL3NetworkToVmNicParam) (*view.VmNicInventoryView, error) {
	resp := view.VmNicInventoryView{}
	if err := cli.Post(fmt.Sprintf("v1/nics/%s/l3-networks/%s", vmNicUuid, l3NetworkUuid), params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateVmNicDriver updates VmNicDriver
func (cli *ZSClient) UpdateVmNicDriver(vmInstanceUuid string, params param.UpdateVmNicDriverParam) (*view.VmNicInventoryView, error) {
	resp := view.VmNicInventoryView{}
	if err := cli.PutWithRespKey("v1/vm-instances", vmInstanceUuid, "", map[string]interface{}{
		"updateVmNicDriver": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetNicQos gets NicQos by uuid
func (cli *ZSClient) GetNicQos(uuid string) (*view.GetNicQosView, error) {
	var resp view.GetNicQosView
	if err := cli.GetWithRespKey("v1/vm-instances", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ChangeVmNicNetwork changes VmNicNetwork
func (cli *ZSClient) ChangeVmNicNetwork(vmNicUuid string, destL3NetworkUuid string, params param.ChangeVmNicNetworkParam) (*view.VmNicInventoryView, error) {
	resp := view.VmNicInventoryView{}
	if err := cli.Post(fmt.Sprintf("v1/vm-instances/nics/%s/l3-networks/%s", vmNicUuid, destL3NetworkUuid), params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AttachVmNicToVm operates on VmNicToVm
func (cli *ZSClient) AttachVmNicToVm(vmInstanceUuid string, vmNicUuid string, params param.AttachVmNicToVmParam) (*view.VmInstanceInventoryView, error) {
	resp := view.VmInstanceInventoryView{}
	if err := cli.Post(fmt.Sprintf("v1/vm-instances/%s/nices/%s", vmInstanceUuid, vmNicUuid), params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ChangeVmNicState changes VmNicState
func (cli *ZSClient) ChangeVmNicState(vmNicUuid string, params param.ChangeVmNicStateParam) (*view.VmInstanceInventoryView, error) {
	resp := view.VmInstanceInventoryView{}
	if err := cli.PutWithRespKey("v1/vm-instances/nics", vmNicUuid, "", map[string]interface{}{
		"changeVmNicState": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
