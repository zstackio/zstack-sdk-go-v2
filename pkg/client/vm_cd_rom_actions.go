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

// CreateVmCdRom creates VmCdRom
func (cli *ZSClient) CreateVmCdRom(params param.CreateVmCdRomParam) (*view.VmCdRomInventoryView, error) {
	resp := view.VmCdRomInventoryView{}
	if err := cli.Post("v1/vm-instances/cdroms", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// DeleteVmCdRom deletes VmCdRom
func (cli *ZSClient) DeleteVmCdRom(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/vm-instances/cdroms", uuid, string(deleteMode))
}
// QueryVmCdRom queries VmCdRom list
func (cli *ZSClient) QueryVmCdRom(params *param.QueryParam) ([]view.VmCdRomInventoryView, error) {
	var resp []view.VmCdRomInventoryView
	return resp, cli.List("v1/vm-instances/cdroms", params, &resp)
}

func (cli *ZSClient) GetVmCdRom(uuid string) (*view.VmCdRomInventoryView, error) {
	var resp view.VmCdRomInventoryView
	if err := cli.Get("v1/vm-instances/cdroms", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageVmCdRom Pagination
func (cli *ZSClient) PageVmCdRom(params *param.QueryParam) ([]view.VmCdRomInventoryView, int, error) {
	var vmCdRoms []view.VmCdRomInventoryView
	total, err := cli.Page("v1/vm-instances/cdroms", params, &vmCdRoms)
	return vmCdRoms, total, err
}
// UpdateVmCdRom updates VmCdRom
func (cli *ZSClient) UpdateVmCdRom(uuid string, params param.UpdateVmCdRomParam) (*view.VmCdRomInventoryView, error) {
	resp := view.VmCdRomInventoryView{}
	if err := cli.PutWithRespKey("v1/vm-instances/cdroms", uuid, "", map[string]interface{}{
		"updateVmCdRom": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetCandidateIsoForAttachingVm gets CandidateIsoForAttachingVm by uuid
func (cli *ZSClient) GetCandidateIsoForAttachingVm(uuid string) (*view.ImageInventoryView, error) {
	var resp view.ImageInventoryView
	if err := cli.GetWithRespKey("v1/vm-instances", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AttachIsoToVmInstance operates on IsoToVmInstance
func (cli *ZSClient) AttachIsoToVmInstance(vmInstanceUuid string, isoUuid string, params param.AttachIsoToVmInstanceParam) (*view.VmInstanceInventoryView, error) {
	resp := view.VmInstanceInventoryView{}
	if err := cli.Post(fmt.Sprintf("v1/vm-instances/%s/iso/%s", vmInstanceUuid, isoUuid), params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetCandidateVmForAttachingIso gets CandidateVmForAttachingIso by uuid
func (cli *ZSClient) GetCandidateVmForAttachingIso(uuid string) (*view.VmInstanceInventoryView, error) {
	var resp view.VmInstanceInventoryView
	if err := cli.GetWithRespKey("v1/images/iso", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DetachIsoFromVmInstance sends the action with its default options.
func (cli *ZSClient) DetachIsoFromVmInstance(uuid string, deleteMode param.DeleteMode) error {
	return cli.DetachIsoFromVmInstanceWithParams(uuid, deleteMode, param.DetachIsoFromVmInstanceParamDetail{})
}

// DetachIsoFromVmInstanceWithParams sends the required route and query parameters for DetachIsoFromVmInstance.
func (cli *ZSClient) DetachIsoFromVmInstanceWithParams(vmInstanceUuid string, deleteMode param.DeleteMode, params param.DetachIsoFromVmInstanceParamDetail) error {
	query := url.Values{}
	if params.IsoUuid != nil {
		query.Set("isoUuid", *params.IsoUuid)
	}
	return cli.deleteActionWithQuery("v1/vm-instances", vmInstanceUuid, "iso", deleteMode, query)
}

// SetVmInstanceDefaultCdRom operates on VmInstanceDefaultCdRom
func (cli *ZSClient) SetVmInstanceDefaultCdRom(vmInstanceUuid string, uuid string, params param.SetVmInstanceDefaultCdRomParam) (*view.VmCdRomInventoryView, error) {
	resp := view.VmCdRomInventoryView{}
	err := cli.PutWithSpec("v1/vm-instances", vmInstanceUuid, fmt.Sprintf("cdroms/%s/actions", uuid), "", params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
